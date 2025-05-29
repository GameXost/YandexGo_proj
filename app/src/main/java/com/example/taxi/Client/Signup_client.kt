package com.example.taxi.Client

import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.widget.Button
import android.widget.TextView
import android.widget.Toast
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import com.example.taxi.R

class Signup_client : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_signup_client)

        // при регистрации запоминается, что ты зарегался
        var memory = getSharedPreferences("Login", Context.MODE_PRIVATE).edit()
        val name: TextView = findViewById(R.id.Sign_up_firstname)
        val secondname: TextView = findViewById(R.id.Sign_up_secondname)
        val email: TextView = findViewById(R.id.Sign_up_email)
        val number: TextView = findViewById(R.id.Sign_up_number)
        val password: TextView = findViewById(R.id.Sign_up_password)
        val button: Button = findViewById(R.id.Sign_up_button)
        var signinText: TextView = findViewById(R.id.Sign_up_to_sign_in)

        // переход на страницу входа, при нажатии на синий текст
        signinText.setOnClickListener {
            startActivity(Intent(this, Signin_client::class.java))
        }

        // Проверки кнопки зарегаться на заполнение полей
            button.setOnClickListener {
                if (name.text.isEmpty()) {
                    Toast.makeText(this, "Заполните поле Имя", Toast.LENGTH_LONG).show()
                } else if (secondname.text.isEmpty()){
                    Toast.makeText(this, "Заполните поле Фамилия", Toast.LENGTH_LONG).show()
                } else if (email.text.contains("@") == false){
                    Toast.makeText(this, "Адрес почты неккоректен", Toast.LENGTH_LONG).show()
                } else if (number.text.isEmpty() || number.text.length != 11) {
                    Toast.makeText(
                        this,
                        "Номер телефона заполнен неккоректно",
                        Toast.LENGTH_LONG
                    ).show()
                } else if (password.text.isEmpty() || password.text.length < 9) {
                    Toast.makeText(
                        this,
                        "Длина пароля должна быть не меньше 8 символов",
                        Toast.LENGTH_LONG
                    ).show()
                } else {
                    // переход на main после регистрации, если бд подключим, сделаем переход на страницу входа
                    memory.putString("Email", email.text.toString()).apply()
                    startActivity(Intent(this, Main_client::class.java))
                }
            }

    }

}
