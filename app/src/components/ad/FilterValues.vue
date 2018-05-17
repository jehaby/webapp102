<template>
<div>
    <label class="label">{{ name }}</label>
    <div class="field has-addons has-addons-centered">

        <div class="control">
            <div class="select is-multiple">
            <p>
                <label v-for="val in values" class="checkbox">
                    <input type="checkbox" :id="val" :value="val" v-model="selected">
                    <!-- TODO: event handling -->
                        {{ val }}
                </label>
            </p>

            </div>
        </div>
    </div>

</div>
</template>

<script>
export default {
  name: 'FilterValues',
  props: [ 'name', 'values', 'chosen' ],
  data () {
    return {
      selected: this.chosen
    }
  },
  watch: {
    chosen () {
      this.selected = this.chosen
    },
    selected: {
      handler (val) {
        console.log('in selected watcher', val)
        this.$emit('update', val)
      },
      deep: true
    }
  }
}
</script>

