package com.example.taxi.Driver

import android.content.Context
import android.os.Bundle
import android.view.View
import android.widget.TextView
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import com.example.taxi.R
import com.google.android.material.bottomsheet.BottomSheetBehavior

class Main_driver : AppCompatActivity()
{
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_main_driver)

        var memory = getSharedPreferences("Login", Context.MODE_PRIVATE)
        memory.edit().putString("LO", "1").apply()
        var emailname : TextView = findViewById(R.id.check)
        emailname.text = memory.getString("Email", "Ты хуесос")

        val bottomSheetBehavior: BottomSheetBehavior<*>?
        val bottomSheet: View = findViewById(R.id.sheet)
        var a = 0
        bottomSheetBehavior = BottomSheetBehavior.from(bottomSheet).apply {
            peekHeight = 20
            this.state=BottomSheetBehavior.STATE_COLLAPSED
        }

    }

    override fun onBackPressed() {

    }
}