package com.example.taxi.Client

import android.content.pm.PackageManager
import android.content.Context
import android.os.Bundle
import android.view.View
import android.widget.TextView
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import com.example.taxi.R
import com.google.android.material.bottomsheet.BottomSheetBehavior
import com.yandex.mapkit.Animation
import com.yandex.mapkit.MapKit
import com.yandex.mapkit.MapKitFactory
import com.yandex.mapkit.geometry.Point
import com.yandex.mapkit.map.CameraPosition
import com.yandex.mapkit.mapview.MapView
import com.yandex.mapkit.map.MapObjectTapListener
import com.yandex.mapkit.search.*
import com.yandex.mapkit.places.panorama.Position
import com.yandex.runtime.image.ImageProvider

private lateinit var mapView: MapView

class Main_client : AppCompatActivity()
{
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()

        // Инициализация карты

        MapKitFactory.setApiKey("988a6ee0-e593-41ab-950a-09d2cd8510a3") // мне похуй
        MapKitFactory.initialize(this)
        setContentView(R.layout.activity_main_client)

        mapView = findViewById(R.id.mapview)
        val mapKit:MapKit = MapKitFactory.getInstance()


        requestLocationPermission()
        var locationLayerMapkit = mapKit.createUserLocationLayer(mapView.mapWindow)
        mapKit.createLocationManager()
        locationLayerMapkit.isVisible = true

        // Конец инициализации карты

        // хуйня которая берёт текст, который кладётся в файле регистрации, можно реализовать как запоминание текущего юзера
        var memory = getSharedPreferences("Login", Context.MODE_PRIVATE)
        memory.edit().putString("LO", "1").apply()
        var emailname : TextView = findViewById(R.id.check)

        // нижняя двигающиеся панелька
        val bottomSheetBehavior: BottomSheetBehavior<*>?
        val bottomSheet: View = findViewById(R.id.sheet)
        var a = 0
        bottomSheetBehavior = BottomSheetBehavior.from(bottomSheet).apply {
            peekHeight = 20
            this.state=BottomSheetBehavior.STATE_COLLAPSED
        }

    }

    override fun onResume() {
        super.onResume()
        val map = mapView.mapWindow.map

        // Установка изначальной позиции на карте
        map.move(
            CameraPosition(
                Point(55.810157, 37.501454),
                /* zoom = */ 15.0f,
                /* azimuth = */ 150.0f,
                /* tilt = */ 15.0f
            ),
            Animation(Animation.Type.SMOOTH, 2f),
            null
        )
    }

    // Функции для запуска и остановки mapkit
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

    /*
    override fun onBackPressed() {

    }

    */

    private fun requestLocationPermission() {
        if (ActivityCompat.checkSelfPermission(this, android.Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED &&
            ActivityCompat.checkSelfPermission(this, android.Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(this, arrayOf(android.Manifest.permission.ACCESS_FINE_LOCATION, android.Manifest.permission.ACCESS_COARSE_LOCATION), 0)
            return
        }
    }
}
