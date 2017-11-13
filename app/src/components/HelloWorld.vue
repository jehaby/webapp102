<template>
  <div class="hello">
    <p> {{ count }}
      <button @click="increment">+</button>

      <button @click="api">api</button>
    </p>
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
  import { getUser } from './../api'

  export default {
    name: 'HelloWorld',
    computed: {
      count () {
        return this.$store.state.count
      }
    },
    methods: {
      increment () {
        this.$store.commit('increment')
      },
      api () {
        getUser()
      }
    },
    data () {
      return {
        msg: 'WTF?'
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
