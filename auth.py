import secrets
import string
import json
import email_validator
import smtplib
import psycopg2

class User:
    def __init__(self, first_name: str, last_name: str, email: str):
        self.id = self.generate_id()
        self.first_name = first_name
        self.last_name = last_name
        self.email = email

    def generate_id(self):
        alphabet = string.ascii_letters + string.digits
        return ''.join(secrets.choice(alphabet) for _ in range(16))

    def to_json(self):  # Егыч тут тож, какой тебе формат удобнее, сделай
        userJson = {'first_name': self.first_name,
                    'last_name': self.last_name, 'email': self.email, 'id': self.id}
        return userJson

    def first_validate(self):
        if any(char.isdigit() for char in self.first_name):
            # (бэк)Я это так чисто написал как заглушки, надо что то поадыкватнее придумать
            if not self.first_name.strip() or not self.first_name.isalpha():
                raise ValueError("Недействительное имя")

        if any(char.isdigit() for char in self.last_name):
            if not self.last_name.strip() or not self.last_name.isalpha():
                raise ValueError("Недействительная фамилия")

        try:
            validated_email = email_validator.validate_email(
                self.email, check_deliverability=True)
            self.email = validated_email.email
        except email_validator.EmailNotValidError:
            raise email_validator.EmailNotValidError(
                f"Недействительный email")
        return True

    def second_validate(self):
        emailvalidator = EmailValidator(self.email)
        emailvalidator.send_verification_email()
        user_code = input()  # (фронт)точка входа для ввода кода потдтверждения
        if emailvalidator.compare_codes(user_code=user_code):
            return True
        else:
            raise KeyError("Неверный код")


# (фронт)было бы неплохо если бы это выводилось красиво а не просто строка с кодом
class EmailValidator:
    def __init__(self, reciever: str):
        self.sender = "email_sender89@mail.ru"
        self.password = "kfgcuQSTttK10Bymbs8B"
        self.reciever = reciever
        self.verification_code = self.verification_code_generator()

    def verification_code_generator(self):
        alphabet = string.digits
        return ''.join(secrets.choice(alphabet) for _ in range(5))

    def send_verification_email(self):
        try:
            server = smtplib.SMTP("smtp.mail.ru", 587)
            server.starttls()
            server.login(self.sender, self.password)
            message = f"Subject: Verification code\nFrom: {self.sender}\nTo: {self.reciever}\n\nThis is your verification code:{self.verification_code}"
            server.sendmail(self.sender, self.reciever, message)
        except Exception as e:
            print(f"Ошибка при отправке: {str(e)}")
        finally:
            server.quit()

    def compare_codes(self, user_code: str):
        if user_code == self.verification_code:
            return True
        else:
            raise KeyError("Неверный код")

# Егыч это тебе ебаться с этим классом, как в базу пихать пользователей. json формат для заглушки(оно даже не работает если что, не юзайте)
class UserRepository:
    def __init__(self):
        ### conn to pg
        self.conn = psycopg2.connect(
            dbname="users",
            user="test",
            password="0000",
            host="localhost",
            port="5432"
        )

        ### cursor 4 req
        self.cur = self.conn.cursor()

    def create_users_table(self):
        query = """
        CREATE TABLE users (id, first_name, last_name, email, phone_number)
        """

    def add_user(self, user: User):
        # if user.first_validate() and user.second_validate():
        #     with open("BASE.json", "w") as UserBase:
        #         json.dump(user.to_json(), UserBase, ensure_ascii=False)   ###хз че с этим делать. я пишу через sql запрос

        ### sql req
        query = """
               INSERT INTO users (id, first_name, last_name, email, phone_number)
               VALUES (%s, %s, %s, %s, %s::jsonb);
               """

        user_data = {
            "id": user.generate_id,
            "first_name": user.first_name,
            "last_name": user.last_name,
            "email": user.email,
            "phone_number": user.phone_number
        }

        self.cur.execute(query, (user_data["user_id"], user_data["first_name"],
                                 user_data["last_name"], user_data["email"],
                                 user_data["phone_number"]))

        self.conn.commit()

    def delete_User(self, user: User):
        self.cur.execute("DELETE FROM users WHERE id = %s", (id, ))

    def list_all(self, user: User):
        self.cur.execute("SELECT * FROM users")
        users = self.cur.fetchall()
        return users

    def find_by_id(self, id: str):
        self.cur.execute("SELECT * FROM users WHERE id = %s", (id, ))

    def close_conn(self):
        self.cur.close()
        self.conn.close()

def main():
    user1 = User("Афелок", "Конченный", "bogdanovmihail129@gmail.com")
    print(user1.first_validate(), user1.second_validate())


if __name__ == "__main__":
    main()