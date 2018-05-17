<template>
<div>
    <label class="label">{{ name }}</label>
    <div class="field has-addons has-addons-centered">
        <p class="control">
            <input v-model="selected.min" class="input" type="number" placeholder="min">
        </p>
    </div>
    <div class="field has-addons has-addons-centered">
        <p class="control">
            <input v-model="selected.max" class="input" type="number" placeholder="max">
        </p>
    </div>
</div>
</template>

<script>
export default {
  name: 'FilterRange',
  props: [ 'name', 'values', 'chosen' ],
  data () {
    return {
      selected: {min: '', max: ''}
    }
  },
  watch: {
    chosen () {
      this.selected = this.chosen
    },
    selected: {
      handler (val) {
        this.$emit('update', val)
      },
      deep: true
    }
  }
}
</script>

