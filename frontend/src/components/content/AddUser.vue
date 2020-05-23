<template>
    <div class="contanier">
        <form>
            <fieldset>
                <legend>Add User</legend>
                <input type="text" id="username" placeholder="Username" v-model="name"><br />
                <input type="text" id="email" placeholder="Email-Adresse" v-model="email"><br />
                <input type="text" id="password" placeholder="Password" v-model="password"><br />
                <button class="primary" @click.prevent="createUser()">Save User</button>
            </fieldset>
        </form>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            email: '',
            name: '',
            password: '',
            errors: [],
        };
    },
    methods: {
        createUser() {
            axios.post('http://192.168.99.102:8081/user', {
                email: this.email,
                name: this.name,
                pw: this.password,
                })
                .then((response) => { console.log(response); })
                .catch((e) => {
                this.errors.push(e);
                });
        },
    },
};
</script>
