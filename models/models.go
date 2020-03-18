package models

// type Repo struct {
//   Id uint `json:"id" gorm:"primary_key"`
//   Owner string `json:"owner"`
//   Name string `json:"name"`
//   Users []User `json:"users" gorm:"many2many:user_repos"`
// }

// type User struct {
//   Id uint `json:"id" gorm:"primary_key"`
//   Name string `json:"name"`
//   Repos []Repo `json:"repos" gorm:"many2many:user_repos"`
// }

// type Commit struct {
//     Id uint `json:"id" gorm:"primary_key"`
//     Hash string `json:"hash"`
//     User User `gorm:"foreignkey:UserId"`
//     UserId uint
//     Repo Repo `gorm:"foreignkey:RepoId"`
//     RepoId uint
// }
