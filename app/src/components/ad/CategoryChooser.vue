<template>
  <div class="columns">
    <div class="column is-12" v-if="done">
      <div class="field is-grouped">
        <p class="control">
          <strong>{{ currentCategoryName }}</strong>
        </p>
        <p class="control">
          <a class="button" v-on:click="change">change</a>
        </p>
      </div>
      
    </div>
    <div v-else v-for="ids,level in show"
         class="category-list column is-4"
    >
      <ul>
        <li
          v-for="id in ids"
          v-on:click="choose(id)"
          
        >
          <p v-bind:class="{ choosen: currentPath.indexOf(id) !== -1 }">
            {{ categories[id].name }}
          </p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import CATEGORIES from '../../graphql/Categories.gql'

export default {
  name: 'CategoryChooser',

  data () {
    return {
      done: false,
      currentPath: [],
      show: [],
      categories: {}
    }
  },
  computed: {
    currentCategoryName () {
      return this.done
        ? this.categories[this.currentPath[this.currentPath.length - 1]].name
        : ''
    }
  },
  async created () {
    // TODO: probably should move to store, use local storage or something for caching

    // TODO: show err message if couldn't load the data
    //   this.$store.dispatch('error', e.message)
  },
  methods: {
    choose (categoryId) {
      categoryId = parseInt(categoryId)
      this.currentPath = this.categories[categoryId].path
      this.show = []
      this.currentPath.forEach((catId, level) => {
        this.show[level] = this.siblings(catId, this.categories)
      })
      this.show[this.currentPath.length] = this.children(categoryId, this.categories)
      if (this.children(categoryId, this.categories).length === 0) {
        this.done = true // leaf category chosen
        this.$emit('chosen', this.currentPath[this.currentPath.length - 1])
      }
    },
    siblings (id, categories) {
      // TODO: move maybe
      id = parseInt(id) // TODO: Types!
      if (categories[id].path.length === 1) {
        return [...new Set(Object.values(categories).map(v => v.path[0]))]
      }
      const parentId = categories[id].path[categories[id].path.indexOf(id) - 1]
      return this.children(parentId, categories)
    },
    children (id, categories) {
      id = parseInt(id) // TODO: Types!
      const res = []
      Object.values(categories).forEach(val => {
        let i = val.path.indexOf(id)
        if (i !== -1 && val.path.length === i + 2) {
          res.push(val.path[i + 1])
        }
      })
      return res // TODO: sort maybe
    },
    change () {
      this.done = false
      this.$emit('chosen', 0) // TODO: use constant for zero category maybe
    }
  },

  apollo: {
    categories: {
      query: CATEGORIES,
      update ({ categories }) {
        // TODO: types!
        var res = categories.reduce(
          (acc, cat) => ({ ...acc, [cat.id]: cat }),
          {}
        )
        this.show = [this.siblings(1, res)]
        return res
      }
    }
  }
}
</script>

<style scoped>
.category-list {
  display: inline;
  border: 1px;
}

.choosen {
  background: lightslategray;
}
</style>
