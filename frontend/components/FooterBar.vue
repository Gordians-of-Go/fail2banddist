<script setup>
import { containerMaxW } from '@/configs/config.js'


const year = new Date().getFullYear()
</script>

<template>
  <footer class="py-2 px-6" :class="containerMaxW">
    <BaseLevel>
      <div class="text-center md:text-left">
        <b>&copy;{{ year }}, <a href="https://app.me/" target="_blank">app.me</a>.</b>
        <slot />
      </div>
      <div class="md:py-2">
        
      </div>
    </BaseLevel>
  </footer>
</template>
