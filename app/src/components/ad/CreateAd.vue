<template>
  <div class="column is-8">

    <h1 class="title"> Create Ad </h1>

    <form>
      
      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">category_choose</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <category-chooser v-on:chosen="categoryChosen"></category-chooser>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">name</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <input v-model="ad.name" class="input" type="text" required>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">description</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <input v-model="ad.description" class="input" type="text" required></input>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">locality</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <textarea v-model="ad.locality" class="input" type="text" required></textarea>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">price</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <input v-model="ad.price" class="input" type="text" required></input>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">weight</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <input v-model="ad.weight" class="input" type="text" required></input>
                  <!-- TODO: currencies -->
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">brand</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <input v-model="ad.brand" class="input" type="text" required></input>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <!-- Left empty for spacing -->
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <button v-on:click="create" :disabled="! ready" class="button is-primary">
                Create
              </button>
            </div>
          </div>
        </div>
      </div>

    </form>
  </div>
</template>

<script>
import { createAd } from '../../api/ad.js'
import CategoryChooser from './CategoryChooser'

export default {
  components: { CategoryChooser },
  name: 'CreateAd',
  data () {
    return {
      ad: {
        name: '',
        description: '',
        category_id: 0
      }
    }
  },
  computed: {
    ready () {
      return (
        this.ad.category_id !== 0 &&
        this.ad.name.length > 5 &&
        this.ad.description.length > 5
      )
    }
  },
  methods: {
    categoryChosen (id) {
      this.ad.category_id = id
      console.log('in create ad', id)
    },
    async create () {
      // TODO: form validation
      let resp = {}
      try {
        console.log('in try component', this.$store.state.auth.jwtToken)
        resp = await createAd({ ...this.ad }, this.$store.state.auth.jwtToken)
      } catch (e) {
        console.log('ad creation failed', e)
        // TODO: better errors
        return this.$store.dispatch('error', 'Ad creation failed')
      }
      console.log('got resp in craeteAD: ', resp)
      this.$router.push('/ads/' + resp.data.ad.uuid)
    }
  }
}
</script>

<style scoped>

</style>
