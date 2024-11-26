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
            value={formStoreData.data.password}
            error={formStoreData.errors.password?.[0] || null}
            id="password"
            type="password"
            label="Password"
    />
    <Input
            wasSubmitted={formStoreData.wasSubmitted}
            value={formStoreData.data.confirmPassword}
            error={formStoreData.errors.confirmPassword?.[0] || null}
            id="confirmPassword"
            type="password"
            label="Confirm Password"
    />
{/if}