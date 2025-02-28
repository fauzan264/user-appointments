<script setup>
import { ref } from 'vue';

const props = defineProps({
    user: Object,
    errors: Object,
})

// Emits to parent
const emit = defineEmits(["register"]);

// checkbox state
const checkbox = ref(true);
</script>

<template>
    <form @submit.prevent="emit('register')">
        <v-row class="d-flex mb-3">
            <v-col cols="12">
                <p v-if="errors" class="text-caption bg-error text-center">{{ errors.errors }}</p>
            </v-col>
            
            <v-col cols="12">
                <v-label class="font-weight-bold mb-1">Name</v-label>
                <v-text-field v-model="user.name" variant="outlined" hide-details color="primary"></v-text-field>
            </v-col>
            <v-col cols="12">
                <v-label class="font-weight-bold mb-1">Username</v-label>
                <v-text-field v-model="user.username" variant="outlined" hide-details color="primary"></v-text-field>
            </v-col>
            <v-col cols="12">
                <v-label class="font-weight-bold mb-1">Password</v-label>
                <v-text-field v-model="user.password" variant="outlined" type="password"  hide-details color="primary"></v-text-field>
            </v-col>
            <v-col cols="12" >
                <v-btn type="submit" color="primary" size="large" block flat>Sign up</v-btn>
            </v-col>        
        </v-row>
    </form>
</template>