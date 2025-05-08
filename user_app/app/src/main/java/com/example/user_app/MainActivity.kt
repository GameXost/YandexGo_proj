package com.example.user_app

import android.content.pm.PackageManager
import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import com.yandex.mapkit.Animation
import com.yandex.mapkit.MapKit
import com.yandex.mapkit.MapKitFactory
import com.yandex.mapkit.geometry.Point
import com.yandex.mapkit.map.CameraPosition
import com.yandex.mapkit.map.MapObjectTapListener
import com.yandex.mapkit.search.*
import com.yandex.mapkit.mapview.MapView
import com.yandex.runtime.image.ImageProvider

private lateinit var mapView: MapView

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        // Инициализация карты
        MapKitFactory.setApiKey("35abb305-18a4-4383-acb4-6ddc53675e92")
        setContentView(R.layout.activity_main)

        MapKitFactory.initialize(this)
        setContentView(R.layout.activity_main)
        mapView = findViewById(R.id.mapview)
        val mapKit:MapKit = MapKitFactory.getInstance()


        requestLocationPermission()
        var locationLayerMapkit = mapKit.createUserLocationLayer(mapView.mapWindow)
        mapKit.createLocationManager()
        locationLayerMapkit.isVisible = true

    }

    private fun requestLocationPermission() {
        if (ActivityCompat.checkSelfPermission(this, android.Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
                ActivityCompat.checkSelfPermission(this, android.Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(this, arrayOf(android.Manifest.permission.ACCESS_FINE_LOCATION, android.Manifest.permission.ACCESS_COARSE_LOCATION), 0)
            return
        }
    }

    override fun onResume() {
        super.onResume()
        val map = mapView.mapWindow.map

        // Установка изначальной позиции на карте
        map.move(
            CameraPosition(
                Point(55.810157, 37.501454),
                /* zoom = */ 10.0f,
                /* azimuth = */ 150.0f,
                /* tilt = */ 30.0f
            ),
            Animation(Animation.Type.SMOOTH, 2f),
            null
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