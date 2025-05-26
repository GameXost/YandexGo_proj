package com.example.taxi

import android.content.Intent
import android.os.Bundle
import android.widget.Button
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import com.example.taxi.Client.Signin_client
import com.example.taxi.Driver.Signin_driver
import com.example.taxi.Driver.Signup_driver

class Activity_choose : AppCompatActivity()
{
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_choose)

        val Client_button : Button = findViewById(R.id.Client_button)
        val Driver_button : Button = findViewById(R.id.Driver_button)

        Client_button.setOnClickListener {
            val intent = Intent(this, Signin_client::class.java)
            startActivity(intent)
        }

        Driver_button.setOnClickListener {
            val intent = Intent(this, Signin_driver::class.java)
            startActivity(intent)
        }

    }
}
