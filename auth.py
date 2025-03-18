import secrets
import string
import email_validator
import smtplib
import psycopg2



class User:
    def __init__(self, first_name: str, last_name: str, email: str, phone_number: str):
        self.id = self.generate_id()
        self.first_name = first_name.strip()
        self.last_name = last_name.strip()
        self.email = email.strip()
        self.phone_number = phone_number.strip()  # без +7
        self.validate_name(self.first_name, 'Имя')
        self.validate_name(self.last_name, 'Фамилия')
        self.validate_email()
        self.validate_phone_number()

    def validate_name(self,name: str, field_name: str):
        if not name:
            raise NameError(f"{field_name} не может быть пустым")
        if not name.isalpha():
            raise NameError(
                f"{field_name} должно содержать только буквы")

    def generate_id(self):
        alphabet = string.ascii_letters + string.digits
        return ''.join(secrets.choice(alphabet) for _ in range(16))
    def validate_email(self):
        try:
            validated_email = email_validator.validate_email(
                self.email, check_deliverability=True)
            self.email = validated_email.email
        except email_validator.EmailNotValidError:
            raise email_validator.EmailNotValidError(
                f"Недействительный email")

    def validate_phone_number(self):
        if len(self.phone_number) != 10:
            raise NameError("Неверный формат телефона")

    def verificate_email(self, user_code: str):
        emailvalidator = EmailValidator(self.email)
        emailvalidator.send_verification_email()
        if not emailvalidator.compare_codes(user_code=user_code):
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


class UserRepository:
    def __init__(self):
        # conn to pg
        self.conn = psycopg2.connect(
            dbname="users",
            user="test",
            password="0000",
            host="localhost",
            port="5432"
        )

        # cursor 4 req
        self.cur = self.conn.cursor()

    def create_users_table(self):
        query = """
        CREATE TABLE users (
        id VARCHAR(16) NOT NULL, 
        first_name VARCHAR(50) NOT NULL, 
        last_name VARCHAR(50) NOT NULL, 
        email VARCHAR(50) NOT NULL, 
        phone_number VARCHAR(10) NOT NULL
        );
        """

    def add_user(self, user: User):
        # sql req
        query = """
               INSERT INTO users (id, first_name, last_name, email, phone_number)
               VALUES (%s, %s, %s, %s, %s::jsonb);
               """

        user_data = {
            "id": user.id,
            "first_name": user.first_name,
            "last_name": user.last_name,
            "email": user.email,
            "phone_number": user.phone_number
        }

        self.cur.execute(query, (user_data["id"], user_data["first_name"],
                                 user_data["last_name"], user_data["email"],
                                 user_data["phone_number"]))

        self.conn.commit()

    def delete_user(self, user: User):
        self.cur.execute("DELETE FROM users WHERE email = %s", (user.email, ))
        self.conn.commit()

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
    user1 = User("Афелок", "Конченный",
                 "bogdanovmihail129@gmail.com", '9937222035')
    Repository = UserRepository()
    Repository.delete_user(user1)

if __name__ == "__main__":
    main()
