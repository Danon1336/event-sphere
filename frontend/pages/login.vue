<template>
    <div
        class="min-h-screen flex items-center justify-center relative"
        :style="{ backgroundColor: '#1A1A2E' }"
    >
        <!-- 🌌 ЗВЁЗДНОЕ НЕБО -->
        <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
            <div
                v-for="i in 50"
                :key="i"
                class="floating-circle"
                :style="randomStyle()"
            />
        </div>

        <!-- 🟢 ОСНОВНОЙ БЛОК -->
        <div
            class="bg-dark-gray backdrop-blur-md p-10 rounded-2xl shadow-xl w-[500px] border border-[#3A3A50] form-container"
        >
            <h2
                class="text-3xl font-semibold text-center text-[#E4E4F0] font-orbitron mb-8 animate-pulse-slow"
            >
                Вход в систему
            </h2>

            <form @submit.prevent="onLogin" class="space-y-6">
                <input
                    v-model="email"
                    type="email"
                    placeholder="Email"
                    class="w-full p-4 rounded-lg bg-[#3A3A50] placeholder-[#E4E4F0]/70 text-[#E4E4F0] focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50 transition-all input-neon"
                />
                <input
                    v-model="password"
                    type="password"
                    placeholder="Пароль"
                    class="w-full p-4 rounded-lg bg-[#3A3A50] placeholder-[#E4E4F0]/70 text-[#E4E4F0] focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50 transition-all input-neon"
                />
                <button
                    type="submit"
                    class="w-full p-4 rounded-lg bg-[#6C63FF] text-[#1A1A2E] font-semibold hover:bg-opacity-90 transition-all ease-in-out shadow-xl hover:shadow-2xl"
                >
                    Войти
                </button>
                <NuxtLink
                    to="/register"
                    class="block text-center text-[#E4E4F0]/70 hover:text-[#E4E4F0] transition-all"
                >
                    Регистрация
                </NuxtLink>
            </form>
        </div>

        <!-- ✅ Уведомление -->
        <div
            v-if="showSuccess"
            class="fixed top-6 right-6 bg-[#6C63FF] text-[#1A1A2E] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-50"
        >
            ✅ Успешный вход!
        </div>

        <!-- ❌ Ошибка входа -->
        <div
            v-if="showError"
            class="fixed top-6 right-6 bg-[#FF6584] text-[#1A1A2E] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-50"
        >
            ❌ Неверный email или пароль
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const email = ref("");
const password = ref("");
const showSuccess = ref(false);
const showError = ref(false);
const csrfToken = ref("");

onMounted(async () => {
    await fetchCSRFToken();
});

const fetchCSRFToken = async () => {
    const res = await fetch("http://localhost:8080/csrf-token", {
        credentials: "include",
    });
    const data = await res.json();
    csrfToken.value = data.csrf_token;
};

const onLogin = async () => {
    try {
        const res = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "X-CSRF-Token": csrfToken.value,
            },
            body: JSON.stringify({
                email: email.value,
                password: password.value,
            }),
            credentials: "include",
        });
        const data = await res.json();
        if (res.ok) {
            showSuccess.value = true;
            setTimeout(() => {
                router.push(`/dashboard/${data.role}`);
            }, 2000);
        } else {
            showError.value = true;
            setTimeout(() => {
                showError.value = false;
            }, 2500);
        }
    } catch (error) {
        console.error("Ошибка:", error);
        showError.value = true;
        setTimeout(() => {
            showError.value = false;
        }, 2500);
    }
};

const randomStyle = () => {
    const size = Math.floor(Math.random() * 60) + 20;
    const left = Math.random() * 100;
    const delay = Math.random() * 10;
    const duration = 10 + Math.random() * 10;
    return {
        width: `${size}px`,
        height: `${size}px`,
        left: `${left}%`,
        animationDelay: `${delay}s`,
        animationDuration: `${duration}s`,
    };
};
</script>

<style scoped>
/* 🌫️ АНИМАЦИИ */
@keyframes floatUp {
    0% {
        transform: translateY(0) scale(1);
        opacity: 0.2;
    }
    50% {
        transform: translateY(-50vh) scale(1.2);
        opacity: 0.5;
    }
    100% {
        transform: translateY(-100vh) scale(1);
        opacity: 0;
    }
}
.floating-circle {
    position: absolute;
    background-color: rgba(108, 99, 255, 0.06);
    border-radius: 9999px;
    animation: floatUp linear infinite;
    filter: blur(1px);
}

/* 🌫️ АНИМАЦИИ */
@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(-20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
@keyframes fadeSlide {
    from {
        opacity: 0;
        transform: translateY(-20px) translateX(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0) translateX(0);
    }
}

.animate-fade-slide {
    animation: fadeSlide 0.5s ease-out;
}

.animate-fadeIn {
    animation: fadeIn 1s ease-out;
}

@keyframes pulseSlow {
    0%,
    100% {
        opacity: 1;
    }
    50% {
        opacity: 0.7;
    }
}
.animate-pulse-slow {
    animation: pulseSlow 3s infinite;
}

/* 🌙 СТИЛИ */
.bg-dark-blue {
    background-color: #1a1a2e;
}

.bg-dark-gray {
    background-color: #2a2a40;
}

.bg-lighter-gray {
    background-color: #3a3a50;
}

.text-light-gray {
    color: #e4e4f0;
}

.text-pink {
    color: #ff6584;
}

.text-purple {
    color: #6c63ff;
}

button {
    background-color: #6c63ff;
    transition: all 0.3s ease;
}

button:hover {
    background-color: #5b54e6;
}

button:focus {
    outline: none;
    box-shadow: 0 0 8px rgba(108, 99, 255, 0.5);
}

button.shadow-xl {
    box-shadow: 0 4px 8px rgba(108, 99, 255, 0.4);
}

button:hover.shadow-2xl {
    box-shadow: 0 8px 16px rgba(108, 99, 255, 0.6);
}

input {
    background-color: #3a3a50;
    color: #e4e4f0;
    border: 1px solid #3a3a50;
    border-radius: 12px;
    padding: 16px;
    transition: all 0.3s ease;
}

input:focus {
    background-color: #4b4b60;
    border-color: #6c63ff;
    box-shadow: 0 0 6px rgba(108, 99, 255, 0.3);
}

input::placeholder {
    color: #e4e4f0;
}

input.input-neon:focus {
    box-shadow: 0 0 12px rgba(108, 99, 255, 1);
    border-color: #6c63ff;
}

a {
    color: #e4e4f0;
    transition: all 0.3s ease;
}

a:hover {
    color: #6c63ff;
}

/* 🌌 ЗВЁЗДЫ */
.floating-circle {
    position: absolute;
    background-color: rgba(108, 99, 255, 0.1);
    border-radius: 50%;
    animation: floatUp linear infinite;
    filter: blur(1px);
}

@keyframes floatUp {
    0% {
        transform: translateY(0) scale(1);
        opacity: 0.3;
    }
    50% {
        transform: translateY(-50vh) scale(1.3);
        opacity: 0.5;
    }
    100% {
        transform: translateY(-100vh) scale(1);
        opacity: 0;
    }
}

/* 🌟 НЕОНОВАЯ ПОЛОСКА ВОКРУГ ФОРМЫ */
.form-container {
    position: relative;
    padding: 20px;
    box-shadow: 0 0 15px rgba(108, 99, 255, 0.6); /* Неоновый эффект */
    border-radius: 12px;
}
</style>
