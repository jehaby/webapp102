<template v-if="message">
  <div>
    <p class="error">{{ message }}</p>
  </div>
</template>

<script>
  export default {
    name: 'error-message',
    props: ['message']
  }
</script>

<style scoped>

</style>
