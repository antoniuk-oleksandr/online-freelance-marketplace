<script lang="ts">
    import Input from "@/common-components/Sign/components/SignInput/Input.svelte";
    import {formStore} from "@/common-components/Sign/stores/form-store.ts";
    import type {FormStore} from "@/types/FormStore.ts";

    let formStoreData = $state<null | FormStore>(null);
    formStore.subscribe((value) => formStoreData = value);
</script>

{#if formStoreData && formStoreData.data && formStoreData.errors}
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.username}
            error={formStoreData.errors.username?.[0] || null}
            id="username"
            type="text"
            label="Username"
    />
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.firstName}
            error={formStoreData.errors.firstName?.[0] || null}
            id="firstName"
            type="text"
            label="First name"
    />
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.surname}
            error={formStoreData.errors.surname?.[0] || null}
            id="surname"
            type="text"
            label="Surname"
    />
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.email}
            error={formStoreData.errors.email?.[0] || null}
            id="email"
            type="text"
            label="Email"
    />
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.password}
            error={formStoreData.errors.password?.[0] || null}
            id="password"
            type="password"
            label="Password"
    />
{/if}