<template>
  <div>
    <Home />
    <!-- <div class="contanier">
        <form>
            <fieldset>
                <legend>Add User</legend>
                <input type="text" id="username" placeholder="Username"><br />
                <input type="text" id="email" placeholder="Email-Adresse"><br />
                <input type="text" id="password" placeholder="Password"><br />
                <button class="primary">Save User</button>
            </fieldset>
        </form>
    </div>  -->
  </div>
</template>

<script>
import Home from './views/Home.vue';

export default {
  components: {
    Home,
  },
};
</script>

<style lang="sass">

</style>
