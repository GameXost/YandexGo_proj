// Top-level build file where you can add configuration options common to all sub-projects/modules.
repositories {
    mavenCentral()
    maven {
        url = uri("http://maven.google.com/")
    }
}

plugins {
    alias(libs.plugins.android.application) apply false
    alias(libs.plugins.kotlin.android) apply false
}
