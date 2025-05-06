package com.example.user_app

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.yandex.mapkit.MapKitFactory
import com.yandex.mapkit.geometry.Point
import com.yandex.mapkit.map.CameraPosition
import com.yandex.mapkit.map.MapObjectTapListener
import com.yandex.mapkit.mapview.MapView
import com.yandex.runtime.image.ImageProvider

private lateinit var mapView: MapView

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        MapKitFactory.setApiKey("35abb305-18a4-4383-acb4-6ddc53675e92")
        setContentView(R.layout.activity_main)

        MapKitFactory.initialize(this)
        setContentView(R.layout.activity_main)
        mapView = findViewById(R.id.mapview)
        val map = mapView.mapWindow.map

        map.move(
            CameraPosition(
                Point(55.751225, 37.62954),
                /* zoom = */ 17.0f,
                /* azimuth = */ 150.0f,
                /* tilt = */ 30.0f
            )
        )
    }

    override fun onStart() {
        super.onStart()
        MapKitFactory.getInstance().onStart()
        mapView.onStart()
    }

    override fun onStop() {
        mapView.onStop()
        MapKitFactory.getInstance().onStop()
        super.onStop()
    }
}