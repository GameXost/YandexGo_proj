package com.example.taxi.Driver

import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.widget.ImageButton
import android.widget.TextView
import android.widget.Toast
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import com.example.taxi.Activity_choose
import com.example.taxi.R

class Signin_driver : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_signin_driver)

        val memory = getSharedPreferences("Login", Context.MODE_PRIVATE)
        if (memory.getString("LO", "0") != "0"){
            startActivity(Intent(this, Main_driver::class.java))
        }
        else{
            var signupText: TextView = findViewById(R.id.Sign_in_to_sign_up)
            var back_to_choose : ImageButton = findViewById(R.id.Back_to_choose)

            back_to_choose.setOnClickListener{
                val intent1 = Intent(this, Activity_choose::class.java)
                startActivity(intent1)
            }
            signupText.setOnClickListener {
                val intent = Intent(this, Signup_driver::class.java)
                startActivity(intent)
            }
        }
    }

}


