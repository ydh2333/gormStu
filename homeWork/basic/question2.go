package basic

// type Account struct {
// 	ID      uint `gorm:"primarykey"`
// 	Balance float64
// }

// type Transaction struct {
// 	ID            uint `gorm:"primarykey"`
// 	FromAccountID uint
// 	ToAccountID   uint
// 	Amount        float64
// }

// func Run(db *gorm.DB) {

// db.AutoMigrate(&Account{})
// db.AutoMigrate(&Transaction{})

// acc := []Account{
// 	{
// 		Balance: 200.00,
// 	},
// 	{
// 		Balance: 50.00,
// 	},
// }
// db.Create(&acc)

// SQL语句如下：
// START TRANSACTION;
// IF (SELECT balance FROM accounts WHERE id = 1) < 100 THEN
// 		ROLLBACK;
// ELSE
// 		UPDATE accounts SET balance = balance - 100 WHERE id = 1;
// 		UPDATE accounts SET balance = balance + 100 WHERE id = 2;
// 		INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (1, 2, 100);
// 		COMMIT;
// END IF;

// 使用gorm触发回滚
// 	tx := db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	acc := Account{
// 		ID: 1,
// 	}
// 	acc2 := Account{
// 		ID: 2,
// 	}
// 	var trans Transaction
// 	tx.First(&acc)
// 	tx.First(&acc2)
// 	fmt.Println("acc.Balance:", acc.Balance)
// 	if acc.Balance < 100 {
// 		fmt.Println("余额不足，触发回滚！")
// 		tx.Rollback()
// 	} else {
// 		acc.Balance -= float64(100)
// 		fmt.Println("acc.Balance after:", acc.Balance)
// 		if err := tx.Model(&acc).Update("balance", acc.Balance).Error; err != nil {
// 			fmt.Println(err)
// 			tx.Rollback()
// 		} else {
// 			trans.FromAccountID = acc.ID
// 		}

// 		acc2.Balance += float64(100)
// 		fmt.Println("acc2.Balance after:", acc2.Balance)
// 		if err := tx.Model(&acc2).Update("balance", acc2.Balance).Error; err != nil {
// 			fmt.Println(err)
// 			tx.Rollback()
// 		} else {
// 			trans.ToAccountID = acc2.ID
// 		}

// 		trans.Amount = float64(100)
// 		if err := tx.Create(&trans).Error; err != nil {
// 			tx.Rollback()
// 		}

// 		fmt.Println("完成转账！")
// 		tx.Commit()
// 	}
// }
