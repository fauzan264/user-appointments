<script setup lang="ts">
definePageMeta({
    layout: 'blank',
    auth: {
        unauthenticatedOnly: true,
        navigateAuthenticatedTo: '/admin/dashboard'
    }
});

// init router
const router = useRouter();

// state user
const user = reactive({
    name: '',
    username: '',
    password: ''
});

// state validation
const errors = ref({});

// destruct signUp from useAuth
const { signUp } = useAuth();

// method register
const register = async () => {
    // define variable
    let name = user.name
    let username = user.username
    let password = user.password
    let preferred_timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

    try {
        // call signUp from useAuth
        await signUp({
            name,
            username,
            password,
            preferred_timezone
        }, {}, {
            preventLoginFlow: true
        });

        // redirect
        router.push({
            path: "/login",
            query: {
                message: "Account has been registered successfully!",
            }
        });
    } catch (error) {
        const err = error as {response?: {_data?:any} }
        // assign state errors with response error data
        if (err.response?._data) {
            errors.value = err.response._data.data;
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
                                <!-- <LayoutFullLogo /> -->
                                <h3>Register</h3>
                            </div>
                            <!-- <div class="text-body-1 text-muted text-center mb-3">Your Social Campaigns</div> -->
                            <AuthRegisterForm v-model:user="user" v-model:errors="errors" @register="register" />
                            <h6 class="text-h6 text-muted font-weight-medium d-flex justify-center align-center mt-3">
                                Already have an Account?
                                <nuxt-link to="/login"
                                    class="text-primary text-decoration-none text-body-1 opacity-1 font-weight-medium pl-2">
                                    Sign In</nuxt-link>
                            </h6>
                        </v-card-item>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>
