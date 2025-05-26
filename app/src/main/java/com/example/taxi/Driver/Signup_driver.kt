package com.example.taxi.Driver

import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.view.View
import android.widget.Button
import android.widget.ImageButton
import android.widget.TextView
import android.widget.Toast
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import com.example.taxi.R

class Signup_driver : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_signup_driver)

        var memory = getSharedPreferences("Login", Context.MODE_PRIVATE).edit()
        val name: TextView = findViewById(R.id.Sign_up_firstname)
        val secondname: TextView = findViewById(R.id.Sign_up_secondname)
        val email: TextView = findViewById(R.id.Sign_up_email)
        val number: TextView = findViewById(R.id.Sign_up_number)
        val password: TextView = findViewById(R.id.Sign_up_password)
        val button: Button = findViewById(R.id.Sign_up_button)
        var signinText: TextView = findViewById(R.id.Sign_up_to_sign_in)
        var driver_license : TextView = findViewById(R.id.Sign_up_license_number)
        var driver_license_date : TextView = findViewById(R.id.Sign_up_license_date)
        var car_number : TextView = findViewById(R.id.Sign_up_number_of_auto)
        var car_model : TextView = findViewById(R.id.Sign_up_marks_of_auto)
        var car_marks : TextView = findViewById(R.id.Sign_up_model_of_auto)
        var car_color : TextView = findViewById(R.id.Sign_up_color_of_auto)
        var button_next : ImageButton = findViewById(R.id.Pole_next)
        var button_back : ImageButton = findViewById(R.id.Pole_back)

        button_next.setOnClickListener{
            name.visibility = View.GONE
            secondname.visibility = View.GONE
            email.visibility = View.GONE
            number.visibility = View.GONE
            password.visibility = View.GONE
            driver_license.visibility = View.VISIBLE
            driver_license_date.visibility = View.VISIBLE
            car_number.visibility = View.VISIBLE
            car_model.visibility = View.VISIBLE
            car_marks.visibility = View.VISIBLE
            car_color.visibility = View.VISIBLE
            button_back.visibility = View.VISIBLE
            button_next.visibility = View.INVISIBLE
        }

        button_back.setOnClickListener{
            name.visibility = View.VISIBLE
            secondname.visibility = View.VISIBLE
            email.visibility = View.VISIBLE
            number.visibility = View.VISIBLE
            password.visibility = View.VISIBLE
            driver_license.visibility = View.GONE
            driver_license_date.visibility = View.GONE
            car_number.visibility = View.GONE
            car_model.visibility = View.GONE
            car_marks.visibility = View.GONE
            car_color.visibility = View.GONE
            button_back.visibility = View.INVISIBLE
            button_next.visibility = View.VISIBLE
        }

        signinText.setOnClickListener {
            startActivity(Intent(this, Signin_driver::class.java))
        }

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
                } else if (driver_license.text.length != 10){
                    Toast.makeText(this, "Номер водительских прав заполнен неккоректно", Toast.LENGTH_LONG).show()
                } else if (driver_license_date.text.length != 8 || driver_license_date.text.contains("-") == false){
                    Toast.makeText(this, "Дата заполнена неккоректно", Toast.LENGTH_LONG).show()
                } else if (car_number.text.length != 7){
                    Toast.makeText(this, "Номера заполнены неккоректно", Toast.LENGTH_LONG).show()
                } else if (car_model.text.isEmpty()){
                    Toast.makeText(this, "Заполните марку автомобиля", Toast.LENGTH_LONG).show()
                } else if (car_marks.text.isEmpty()){
                    Toast.makeText(this, "Заполните модель автомобиля", Toast.LENGTH_LONG).show()
                } else if (car_color.text.isEmpty()){
                    Toast.makeText(this, "Заполните цвет автомобиля", Toast.LENGTH_LONG).show()
                } else {
                    memory.putString("Email", email.text.toString()).apply()
                    startActivity(Intent(this, Main_driver::class.java))
                }
            }

    }

}
