<template>

    <div class="control">
        <div class="select">
            <select v-model="selected">
                <option v-for="locality in localities" v-bind:value="locality.id">
                    {{ locality.name }}
                </option>
            </select>
        </div>
    </div>

</template>

<script>

// TODO: get localities from backend
export default {
  name: 'LocalityChooser',
  data () {
    return {
      localities: [
          {},
          {id: 1, name: 'moscow'},
          {id: 2, name: 'spb'}
      ],
      selected: ''
    }
  },
  watch: {
    selected (chosen) {
      this.$emit('chosen', chosen)
    }
  }
}
</script>
