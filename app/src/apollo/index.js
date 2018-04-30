import { ApolloClient } from 'apollo-client'
import { HttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import VueApollo from 'vue-apollo'
import Vue from 'vue'

Vue.use(VueApollo)

const httpLink = new HttpLink({
    // You should use an absolute URL here
    // TODO: get from config
  uri: 'http://localhost:8899/query'
})

// Create the apollo client
const apolloClient = new ApolloClient({
  // TODO: read docs
  link: httpLink,
  cache: new InMemoryCache(),
  connectToDevTools: true
})

const apolloProvider = new VueApollo({
  defaultClient: apolloClient
})

export default apolloProvider
