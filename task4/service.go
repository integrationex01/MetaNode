package task4

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ## 3. 用户认证与授权
// - 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
// - 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。

func RegisterService(user User, db *gorm.DB) error {
	encPassword, err := EncryptPassword(user.Password)
	user.Password = encPassword
	if err != nil {
		return err
	}
	return db.Create(&user).Error
}

func EncryptPassword(password string) (string, error) {
	// Implement password encryption logic here (e.g., using bcrypt)
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to encrypt password: %v", err)
		return "", err
	}
	return string(bcryptedPassword), nil
}

func LoginService(user User, db *gorm.DB) (string, error) {
	var foundUser User
	err := db.Where("username = ?", user.Username).First(&foundUser).Error
	if err != nil {
		return "", err
	}
	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// JWT generation logic would go here
	return GenerateJWT(user)
}

func CreatePostService(post Post, db *gorm.DB) error {
	return db.Create(&post).Error
}

func GetPostsService(db *gorm.DB) ([]Post, error) {
	posts := []Post{}
	err := db.Model(&Post{}).
		Omit("content").
		Preload("User").
		Find(&posts).Error
	return posts, err
}

func GetPostByIDService(postID uint, db *gorm.DB) (Post, error) {
	var post Post
	err := db.Preload("User").First(&post, postID).Error
	return post, err
}

func UpdatePostService(post Post, userID uint, db *gorm.DB) error {
	// 验证文章作者
	var existingPost Post
	if err := db.First(&existingPost, post.ID).Error; err != nil {
		return err
	}
	if existingPost.UserID != userID {
		return fmt.Errorf("unauthorized: not the author of the post")
	}
	return db.Model(&post).Updates(post).Error
}

// DeletePostService 删除文章，需要验证用户是否为作者
func DeletePostService(postID uint, userID uint, db *gorm.DB) error {
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		return err
	}
	if post.UserID != userID {
		return fmt.Errorf("unauthorized: not the author of the post")
	}
	return db.Delete(&post).Error
}

// CreateCommentService 创建评论
func CreateCommentService(comment Comment, db *gorm.DB) error {
	return db.Create(&comment).Error
}

// GetPostCommentsService 获取文章的所有评论
func GetPostCommentsService(postID uint, db *gorm.DB) ([]Comment, error) {
	var comments []Comment
	err := db.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
