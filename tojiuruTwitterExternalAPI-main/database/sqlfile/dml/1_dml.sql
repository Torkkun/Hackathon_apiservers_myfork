PRAGMA foreign_keys=true;

INSERT INTO user(user_id, name, password, accesstoken, secrettoken) VALUES("Test UserID", "Test User", "30e8a49c277760e91f164a67c88842a9b5e21ba9", "none_access_token", "none_secrettoken");

INSERT INTO user(user_id, name, password) VALUES("1001", "Owner1", "pass");
INSERT INTO user(user_id, name, password) VALUES("1002", "Owner2", "pass");

INSERT INTO user(user_id, name, password) VALUES("0001", "User1", "pass");
INSERT INTO user(user_id, name, password) VALUES("0002", "User2", "pass");

INSERT INTO examination(message_id, message, people_num, user_id, deadline) VALUES("0000", "test message1", 0, "1001", "2099-01-01 01:01:01");
INSERT INTO examination(message_id, message, people_num, user_id, deadline) VALUES("0001", "test message2", 0, "1002", "2099-02-02 02:02:02");

INSERT OR IGNORE INTO replyuser(message_id, reply_id, from_user_id) VALUES("0000","1111","0002");

INSERT INTO replymessage(reply_id, reply_message_id, reply_text, user_id) values ("1111","AAAA","reply test message","0002");

-- テスト用のメモ
-- select message_id, message, people_num, examination.user_id, user.name, deadline, examination.created_at from examination inner join user on user.user_id = examination.user_id;