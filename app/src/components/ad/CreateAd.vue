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
              <textarea v-model="ad.description" class="textarea" type="text" required></textarea>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">condition</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <div class="select">
                <select v-model="ad.condition" name="" id="">
                  <option value="USED">used</option>
                  <option value="USED_LIKE_NEW">used_like_new</option>
                  <option value="NEW">new</option>
                  <option value="MALFUNCTIONED">malfunctioned</option>
                </select>
              </div>
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
              <input v-model="ad.price" class="input" type="number" required>
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
              <input v-model="ad.weight" class="input" type="number">
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">brand</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <brand-chooser v-on:chosen="brandChosen"></brand-chooser>
            </div>
          </div>
        </div>
      </div>

      <div class="field is-horizontal">
        <div class="field-label">
          <label class="label">contacts</label>
        </div>
        <div class="field-body">
          <div class="field is-narrow">
            <div class="control">
              <contacts-chooser v-on:chosen="contactsChosen"></contacts-chooser>
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
              <locality-chooser v-on:chosen="localityChosen"></locality-chooser>
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
import CategoryChooser from './CategoryChooser'
import BrandChooser from './BrandChooser'
import LocalityChooser from './LocalityChooser'
import ContactsChooser from './ContactsChooser'
import AD_CREATE from './../../graphql/AdCreate.gql'

export default {
  components: { CategoryChooser, LocalityChooser, BrandChooser, ContactsChooser },
  name: 'CreateAd',
  data () {
    return {
      ad: {
        // name: '',
        // description: '',
        // category_id: 0,
        // currency: 'RUB'

        name: 'test',
        description: 'test descrip',
        categoryId: 0,
        currency: 'RUB',
        localityId: 1,
        condition: 'USED',
        price: 300,
        weight: 1
      }
    }
  },
  computed: {
    ready () {
      return (
        this.ad.categoryId !== 0 &&
        this.ad.name.length > 5 &&
        this.ad.description.length > 5 &&
        this.ad.localityId > 0 &&
        this.ad.condition !== ''
      )
    }
  },
  methods: {
    categoryChosen (id) { this.ad.categoryId = id.toString() },
    localityChosen (id) { this.ad.localityId = id.toString() },
    brandChosen (id) { this.ad.brandId = id.toString() },
    contactsChosen (c) {
      this.ad.phoneUUID = c
    },
    async create () {
      // TODO: form validation

      let ad = this.ad
      ad.price = parseInt(ad.price * 100)
      ad.userUUID = 'e12087ab-23b9-4d97-8b61-e7016e4e956b'
      console.log('in async create', {...ad})
      try {
        let resp = await this.$apollo.mutate({
          mutation: AD_CREATE,
          variables: {'input': ad}
        // update: (store, { data: { adCreate } }) => {

        // }
        })
        console.log('resp ok', resp)
        this.$router.push('/ads/' + resp.data.ad.uuid)
      } catch (e) {
        console.log('resp err', e)
        return this.$store.dispatch('error', 'Ad creation failed')
      }
        // console.log('in try component', this.$store.state.auth.jwtToken)
        // TODO: better errors
        // return this.$store.dispatch('error', 'Ad creation failed')
      //
    }
  }
}
</script>

<style scoped>

</style>
