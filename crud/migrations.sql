CREATE TABLE IF NOT EXISTS ITEMS (ID VARCHAR(36) PRIMARY KEY, NAME TEXT NOT NULL, ACTIVE INTEGER);

INSERT INTO ITEMS(ID, NAME, ACTIVE)
SELECT '8b933c84-ee60-45a1-848d-428ad3259e2b', 'Full Self Driving (FSD)', 1
WHERE
NOT EXISTS (
SELECT ID FROM ITEMS WHERE ID = '8b933c84-ee60-45a1-848d-428ad3259e2b'
);

INSERT INTO ITEMS(ID, NAME, ACTIVE)
SELECT 'd660b9b2-0406-46d6-9efe-b40b4cca59fc', 'Sentry Mode', 1
WHERE
NOT EXISTS (
SELECT ID FROM ITEMS WHERE ID = 'd660b9b2-0406-46d6-9efe-b40b4cca59fc'
);