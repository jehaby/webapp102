<template>
  <div class="field-body">
    <div class="field is-narrow">
      <div class="control">
        <label class="radio">
          <input type="radio" v-model="newvalue" v-bind:value="true">
          Yes
        </label>
        <label class="radio">
          <input type="radio" v-model="newvalue" v-bind:value="false">
          No
        </label>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdPropertyBool',
  data () {
    return {
      newvalue: this.value
    }
  },
  props: [
    'value'
  ],
  watch: {
    newvalue: function (chosen) {
      this.$emit('input', chosen)
    }
  }

}
</script>

