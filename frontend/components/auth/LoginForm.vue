<script setup>
import { ref } from 'vue';

const props = defineProps({
    successMessage: String,
    errors: Object,
    username: String,
    password: String
});

// Emits to parent
const emit = defineEmits(["update:username", "update:password","login"]);

const checkbox = ref(true);
</script>

<template>
    <form @submit.prevent="()=> { console.log('Login trigerred');emit('login');}">
        <v-row class="d-flex mb-3">
            <v-col cols="12">
                <p v-if="successMessage" class="text-center bg-success">{{ successMessage }}</p>
                <p v-if="errors" class="text-center bg-error">{{ errors.errors }}</p>
            </v-col>
            <v-col cols="12">
                <v-label class="font-weight-bold mb-1">Username</v-label>
                <v-text-field
                    :model-value="props.username"
                    @update:model-value="emit('update:username', $event)"
                    variant="outlined"
                    hide-details
                    color="primary"
                ></v-text-field>
            </v-col>
            <v-col cols="12">
                <v-label class="font-weight-bold mb-1">Password</v-label>
                <v-text-field
                    :model-value="props.password"
                    @update:model-value="emit('update:password', $event)"
                    variant="outlined"
                    type="password" 
                    hide-details
                    color="primary"
                ></v-text-field>
            </v-col>
            <v-col cols="12" class="pt-0">
                <div class="d-flex flex-wrap align-center ml-n2">
                    <v-checkbox v-model="checkbox"  color="primary" hide-details>
                        <template v-slot:label class="text-body-1">Remember this Device</template>
                    </v-checkbox>
                    <div class="ml-sm-auto">
                        <NuxtLink to="/"
                            class="text-primary text-decoration-none text-body-1 opacity-1 font-weight-medium">Forgot
                            Password ?</NuxtLink>
                    </div>
                </div>
            </v-col>
            <v-col cols="12" class="pt-0">
                <v-btn type="submit" color="primary" size="large" block flat>Sign in</v-btn>
            </v-col>
        </v-row>
    </form>
</template>
