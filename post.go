package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const(
	titleSeperator = "Title: "
	descriptionSeperator = "Description: "
	TagSeperator = "Tags: "
)

type Post struct{
	Title, Description string
	Tags []string
	Body string 
}

func NewPost(postFile io.Reader)(Post,error){
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string)string{
		scanner.Scan()
		sanitzed := strings.TrimSpace(scanner.Text())
		text := strings.TrimPrefix(sanitzed,tagName)
		return strings.TrimSpace(text)
	}

	praseTags := func()[]string{
		tagsLine := readMetaLine(TagSeperator)
		tags := []string{}
		for tag := range strings.SplitSeq(tagsLine, ","){
			tags = append(tags, strings.TrimSpace(tag))
		}
		return tags 
	}

	title := readMetaLine(titleSeperator)
	desc := readMetaLine(descriptionSeperator)
	tags := praseTags()
	body := getBody(scanner)


	post:=Post{
		Title: title,
		Description: desc,
		Tags: tags,
		Body: body,
	}
	return post, nil
}

func getBody(scanner *bufio.Scanner) (body string) {
	scanner.Scan() // ignore --- line

	buffer := bytes.Buffer{}
	for scanner.Scan(){
		fmt.Fprintln(&buffer,scanner.Text())
	}

	body= strings.TrimSpace(buffer.String())

	return 
}