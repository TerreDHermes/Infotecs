CREATE TABLE transactions (
                              id SERIAL PRIMARY KEY,
                              time TIMESTAMPTZ DEFAULT now(),
                              wallet_from_id INT,
                              wallet_to_id INT,
                              amount DECIMAL(10,2),
                              FOREIGN KEY (wallet_from_id) REFERENCES wallets(id),
                              FOREIGN KEY (wallet_to_id) REFERENCES wallets(id)
);