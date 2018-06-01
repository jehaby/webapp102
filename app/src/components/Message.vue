<template v-if="message.text">
  <div>
    <p v-bind:class="message.type">{{ message.text }}</p>
  </div>
</template>

<script>
  export default {
    name: 'message',
    props: [
      'message'
    ]
  }
</script>

<style scoped>

</style>
