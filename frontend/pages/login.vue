<script setup lang="ts">
definePageMeta({
  layout: "blank",
  auth: {
    unauthenticatedOnly: true,
    navigateAuthenticatedTo: "/admin/dashboard",
  },
});

const route = useRoute();
const successMessage = ref(route.query.message as string || "");

// state validation
const username = ref('');
const password = ref('');
const errors = ref({});

// destruct signIn from useAuth
const { signIn } = useAuth();

const login = async () => {
    try {
        // call signIn from useAuth
        await signIn({
            username: username.value,
            password: password.value,
        }, {
            callbackUrl: '/admin/dashboard',
            external: false
        });

    } catch (error) {

        const err = error as { response?: { _data?: any } };
        if (err.response?._data) {
            errors.value = err.response._data?.data || { message: err.response._data?.message || "Unknown error" };
        } else {
            errors.value = { message: "An unexpected error occurred." };
        }
    }
}
</script>
<template>
    <div class="authentication">
        <v-container fluid class="pa-3">
            <v-row class="h-100vh d-flex justify-center align-center">
                <v-col cols="12" lg="4" xl="3" class="d-flex align-center">
                    <v-card rounded="md" elevation="10" class="px-sm-1 px-0 withbg mx-auto" max-width="500">
                        <v-card-item class="pa-sm-8">
                            <div class="d-flex justify-center py-4">
                                <h3>Login</h3>
                            </div>
                            <!-- <div class="text-body-1 text-muted text-center mb-3">Your Social Campaigns</div> -->
                            <AuthLoginForm v-model="successMessage" v-model:username="username" v-model:password="password" v-model:errors="errors" @login="login" />
                            <h6 class="text-h6 text-muted font-weight-medium d-flex justify-center align-center mt-3">
                                New to Modernize?
                                <nuxt-link to="/register"
                                    class="text-primary text-decoration-none text-body-1 opacity-1 font-weight-medium pl-2">
                                    Create an account</nuxt-link>
                            </h6>
                        </v-card-item>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>
