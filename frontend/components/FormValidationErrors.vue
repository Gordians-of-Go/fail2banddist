<script setup>
const errors = ref({});

const hasErrors = ref(false);
</script>

<template>
  <NotificationBarInCard v-if="hasErrors" color="danger">
    <b>Whoops! Something went wrong.</b>
    <span v-for="(error, key) in errors" :key="key">{{ error }}</span>
  </NotificationBarInCard>
</template>
