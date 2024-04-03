<script setup>
import { mdiAccount, mdiAsterisk, mdiHomeOutline, mdiAccountPlus, mdiLockQuestion } from '@mdi/js'


const props = defineProps({
  canResetPassword: Boolean,
  status: {
    type: String,
    default: null
  }
})

const form = reactive({
  email: '',
  password: '',
  remember: false,
  processing: false
})

const submit = () => {
 console.log(form)
}


</script>

<template>
  <NuxtLayout>

    <SectionFullScreen v-slot="{ cardClass }" bg="blue">
      <CardBox :class="cardClass" is-form @submit.prevent="submit">
        <FormValidationErrors />

        <NotificationBarInCard v-if="status" color="info">
          {{ status }}
        </NotificationBarInCard>

        <FormField label="Email" label-for="email" help="Please enter your email">
          <FormControl v-model="form.email" :icon="mdiAccount" id="email" autocomplete="email" type="email" required />
        </FormField>

        <FormField label="Password" label-for="password" help="Please enter your password">
          <FormControl v-model="form.password" :icon="mdiAsterisk" type="password" id="password"
            autocomplete="current-password" required />
        </FormField>

        <BaseDivider />

        <BaseLevel>
          <BaseButtons>
            <BaseButton type="submit" color="info" label="Login" :class="{ 'opacity-25': form.processing }"
              :disabled="form.processing" />
          </BaseButtons>
        </BaseLevel>
      </CardBox>
    </SectionFullScreen>
  </NuxtLayout>
</template>
