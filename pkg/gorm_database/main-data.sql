INSERT INTO users (username, email, password, is_admin)
VALUES
    ('admin', 'admin@example.com', '$2a$10$wDloyIFXKydepjTr3/yAY.HpkeRYRFNADs4Ce/a04sTEqb6SYsz.2', true),
    ('user1', 'user1@example.com', '$2a$10$wDloyIFXKydepjTr3/yAY.HpkeRYRFNADs4Ce/a04sTEqb6SYsz.2', false),
    ('user2', 'user2@example.com', '$2a$10$wDloyIFXKydepjTr3/yAY.HpkeRYRFNADs4Ce/a04sTEqb6SYsz.2', false);

INSERT INTO profiles (user_id, phone_number, authentication_level, blocked_level, balance, is_premium)
VALUES
    (1, '1234567890', 0, 0, 5000000000, true),
    (2, '9876543210', 0, 0, 5000000000, false),
    (3, '9876543210', 1, 1, 5000000000, false);

INSERT INTO banking_info (user_id, bank_name, account_number, card_number, expire_date, cvv2)
VALUES
    (2, 'saman', '123456', '654321', '12/13', '123'),
    (2, 'sepah', '12345678', '87654321', '14/15', '456');

INSERT INTO crypto (name, symbol, current_price, buy_fee, sell_fee)
VALUES
    ('Bitcoin', 'BTC', 500, 515, 485),
    ('Ethereum', 'ETH', 300, 313, 387),
    ('Litecoin', 'LTC', 100, 111, 89);


INSERT INTO open_trade (user_id, crypto_id, amount, buy_fee, stop_loss, take_profit)
VALUES
    (1, 1, 10, 515, 0, 0),
    (1, 1, 10, 515, 450, 600),
    (1, 1, 10, 515, 350, 700);

INSERT INTO closed_trade (user_id, crypto_id, amount, buy_fee, sell_fee, profit)
VALUES
    (1, 1, 10, 300, 400, 1000 ),
    (1, 1, 10, 500, 600, 1000 ),
    (1, 1, 10, 600, 500, -1000);
