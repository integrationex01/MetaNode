package task3

import (
	"gorm.io/gorm"
)

// ## SQL语句练习
// ### 题目1：基本CRUD操作
// - 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
//   - 要求 ：
//     - 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//     - 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//     - 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//     - 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

func BaseGormOperations(db *gorm.DB) {
	// 在这里编写使用 GORM 进行数据库操作的代码
	// 创建数据表
	db.AutoMigrate(&Student{})
	// 插入新记录
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	db.Create(&Student{Name: "李四", Age: 17, Grade: "二年级"})
	db.Create(&Student{Name: "赵六", Age: 24, Grade: "五年级"})
	db.Create(&Student{Name: "王五", Age: 14, Grade: "一年级"})
	// 查询年龄大于18岁的学生
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	// 更新姓名为"张三"的学生年级
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	// 删除年龄小于15岁的学生记录
	db.Where("age < ?", 15).Delete(&Student{})
}

// ### 题目2：事务语句
// - 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//   - 要求 ：
//   - 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func TransactionGormOperations(db *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		// 查询转出账户余额(即账户是否存在)
		if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
			return err
		}
		// 检查余额是否足够
		if fromAccount.Balance < amount {
			return gorm.ErrInvalidTransaction
		}
		// 查询转入账户是否存在
		if err := tx.First(&toAccount, toAccountID).Error; err != nil {
			return err
		}
		// 扣除转出账户余额
		fromAccount.Balance -= amount
		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}
		// 增加转入账户余额
		toAccount.Balance += amount
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}
		// 记录转账信息
		transaction := Transaction{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}
