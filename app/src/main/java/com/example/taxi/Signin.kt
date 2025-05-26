package com.example.taxi

import android.content.Context
import android.content.Intent
import android.content.SharedPreferences
import android.os.Bundle
import android.widget.TextView
import android.widget.Toast
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.datastore.core.DataStore
import java.util.prefs.Preferences

class Signin : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_signin)

        val memory = getSharedPreferences("Login", Context.MODE_PRIVATE)
        if (memory.getString("LO", "0") != "0"){
            startActivity(Intent(this, MainActivity::class.java))
        }
        else{
            var signupText: TextView = findViewById(R.id.Sign_in_to_sign_up)
            signupText.setOnClickListener {
                try {
                    val intent = Intent(this, Signup::class.java)
                    startActivity(intent)
                    // finish() вызывайте только если нужно закрыть текущую активность
                } catch (e: Exception) {
                    e.printStackTrace()
                    Toast.makeText(this, "Ошибка перехода: ${e.message}", Toast.LENGTH_SHORT).show()
                }
            }
        }
    }

}


