SELECT c.id, с.title user u
JOIN chat_users cu ON cu.user_id = u.id
JOIN chat c ON c.id = cu.chat_id
WHERE u.name='Вася';


























