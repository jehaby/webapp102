<template>
<div v-model="selected" class="control">
    <div class="select is-multiple">

    <p>
        <label v-for="condition in conditions" class="checkbox">
            <input type="checkbox">
            <!-- TODO: event handling -->
                {{ condition.name }}
        </label>
    </p>

    </div>
</div>
</template>

<script>
// TODO: get localities from backend
export default {
  name: 'ConditionChooser',
  data () {
    return {
      conditions: [
        {'name': 'used', 'value': 'USED'},
        {'name': 'used_like_new', 'value': 'USED_LIKE_NEW'},
        {'name': 'new', 'value': 'NEW'},
        {'name': 'malfunctioned', 'value': 'MALFUNCTIONED'}
      ],
      selected: ''
    }
  },
  watch: {
    selected (chosen) {
      console.log('in condition choser', chosen)
      this.$emit('chosen', chosen)
    }
  }
}
</script>
