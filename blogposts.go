package blogposts

import (
	"io/fs"
)


func NewPostsFromFS(filesystem fs.FS)([]Post,error){
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir{
		post, err := getPost(filesystem, f.Name())
		if err!=nil{
			return nil, err
		}
		posts = append(posts, post)	
	}
	return posts,nil
}

func getPost(fs fs.FS, fileName string)(Post, error){
	postFile, err := fs.Open(fileName)
	if err!=nil {
		return Post{},err
	}
	defer postFile.Close()

	return NewPost(postFile)
}
