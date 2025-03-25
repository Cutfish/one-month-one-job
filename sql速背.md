```sql
-- 1. 创建一个用户表
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,  -- 主键，自增
    username VARCHAR(50) NOT NULL UNIQUE,    -- 用户名，唯一
    email VARCHAR(100) NOT NULL,
    age INT CHECK (age >= 0),                -- 约束：年龄必须非负
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- 创建时间，默认当前时间
) [ENGINE=InnoDB];

-- 2. 创建一个订单表
CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    amount DECIMAL(10,2) NOT NULL DEFAULT 0.00, -- 订单金额，默认0.00
    order_date DATE NOT NULL,
    status ENUM('pending', 'shipped', 'delivered', 'cancelled') NOT NULL DEFAULT 'pending',
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE -- 外键，用户被删除则级联删除订单
) [ENGINE=InnoDB];

-- 3. 插入数据
INSERT INTO users (username, email, age) VALUES 
('Alice', 'alice@example.com', 25),
('Bob', 'bob@example.com', 30),
('Charlie', 'charlie@example.com', 22);

INSERT INTO orders (user_id, amount, order_date, status) VALUES 
(1, 100.50, '2024-03-01', 'shipped'),
(2, 200.75, '2024-03-02', 'pending'),
(3, 150.00, '2024-03-03', 'delivered');

-- 4. 选择所有订单信息，关联用户
SELECT o.order_id, u.username, o.amount, o.order_date, o.status 
FROM orders o
JOIN users u ON o.user_id = u.user_id
[WHERE o.status = 'shipped']  -- [可选：筛选已发货订单]
[ORDER BY o.order_date DESC]; -- [可选：按日期降序排序]

-- 5. 统计不同订单状态的订单数
SELECT status, COUNT(*) AS order_count
FROM orders
GROUP BY status
[HAVING COUNT(*) > 1]  -- [可选：筛选订单数大于1的状态]
[ORDER BY order_count DESC]; -- [可选：按订单数降序排序]

-- 6. 计算每个用户的订单总金额
SELECT u.username, SUM(o.amount) AS total_spent
FROM users u
LEFT JOIN orders o ON u.user_id = o.user_id
GROUP BY u.user_id
[HAVING total_spent > 100]  -- [可选：只显示订单总额大于100的用户]
[ORDER BY total_spent DESC]; -- [可选：按消费总额降序]

-- 7. 更新订单状态
UPDATE orders
SET status = 'delivered'
WHERE order_id = 1;

-- 8. 删除未付款订单
DELETE FROM orders
WHERE status = 'pending' 
[AND order_date < '2024-03-02']; -- [可选：只删除特定日期前的]

-- 9. 使用 UNION 合并查询结果
(SELECT username AS name, email FROM users)
UNION
(SELECT 'Guest' AS name, 'guest@example.com' AS email)
[ORDER BY name ASC]; -- [可选：按名字排序]

-- 10. 使用子查询查找消费最高的用户
SELECT username, total_spent 
FROM (SELECT u.username, SUM(o.amount) AS total_spent
      FROM users u
      LEFT JOIN orders o ON u.user_id = o.user_id
      GROUP BY u.user_id) AS user_spending
WHERE total_spent = (SELECT MAX(SUM(amount)) FROM orders GROUP BY user_id);

```