package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	log "github.com/s00500/env_logger"
)

func main() {
	log.EnableLineNumbers()
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Use: util [command]\n\r")
		fmt.Println("Commands:")
		fmt.Println("    remove - Remove a selected tag\n\r")
		os.Exit(1)
	}

	// Get the output
	switch args[1] {
	case "remove":
		tags := getTags()

		selectedTag := promptSelect("Select tag", tags...)

		log.Info(selectedTag)
		removeTag(selectedTag)
	default:
	}
}

func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		log.Warnf("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

func getTags() []string {
	cmdTag := exec.Command("git", "tag", "-l")
	outputTag, err := cmdTag.Output()
	log.Should(err)
	tags := string(outputTag)
	return splitTags(tags)
}

func splitTags(tags string) []string {
	taglist := strings.Split(tags, "\n")
	taglist = reverseStringArray(taglist)
	if taglist[0] == "" {
		taglist = taglist[1:]
	}
	return taglist
}

func removeTag(tag string) {
	tag = strings.TrimSpace(tag)
	log.Info("Remove tag: ", tag)
	cmdLocalTag := exec.Command("git", "tag", "-d", tag)
	outputLocalTag, err := cmdLocalTag.CombinedOutput()
	log.Should(err)
	log.Println(string(outputLocalTag))

	cmdOriginTag := exec.Command("git", "push", "origin", "--delete", tag)
	outputOriginTag, err := cmdOriginTag.Output()
	log.Should(err)
	log.Println(string(outputOriginTag))
}

func reverseStringArray(arr []string) []string {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func promptSelect(label string, answers ...string) string {
	result := ""
	prompt := &promptui.Select{
		Label: label,
		Items: answers,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
