<template>
  <div class="field-body">
    <div class="field is-narrow">
      <div class="control">
        RANGE!
        <!-- TODO: implement -->
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdPropertyRange',
  data () {
    return {
      newvalue: this.value
    }
  },
  props: [
    'values',
    'value'
  ],
  watch: {
    newvalue: function (chosen) {
      this.$emit('input', chosen)
    }
  }

}
</script>

