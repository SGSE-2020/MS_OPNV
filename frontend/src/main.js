import firebase from 'firebase';
import Vue from 'vue';
import App from './App.vue';
import router from './router';

Vue.config.productionTip = false;

let app = '';
const config = {
  apiKey: 'AIzaSyBvTg0_QrhEvQ9UeZPH8--E2JZ55KA_u_c',
  authDomain: 'smart-city-ss2020.firebaseapp.com',
  databaseURL: 'https://smart-city-ss2020.firebaseio.com',
  projectId: 'smart-city-ss2020',
  storageBucket: 'smart-city-ss2020.appspot.com',
  messagingSenderId: '957240233717',
};

firebase.initializeApp(config);

// new Vue({
//   router,
//   render: (h) => h(App),
// }).$mount('#app');

firebase.auth().onAuthStateChanged(() => {
  if (!app) {
    /* eslint-disable no-new */
    app = new Vue({
      router,
      render: (h) => h(App),
    }).$mount('#app');
  }
});
