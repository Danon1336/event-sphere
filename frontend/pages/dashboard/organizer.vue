<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />

    <div class="container mx-auto p-4" style="padding-top: 80px">
      <!-- Profile Card -->
      <!--
      <div class="flex flex-col md:flex-row gap-6 mb-6">
        <div class="p-6 text-white">
          <h1 class="text-2xl font-bold mb-4">Личный кабинет организатора</h1>

          <div v-if="user">
            <p>Добро пожаловать, {{ user.name }}</p>
            <p><strong>Email:</strong> {{ user.email }}</p>
            <p><strong>Роль:</strong> {{ user.role }}</p>
            <button
              class="mt-4 bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded"
              @click="logout"
            >
              Выйти
            </button>
          </div>
          <div v-else>
            <p>Загрузка данных пользователя...</p>
          </div>
        </div>

        <!-- Stats Cards (без изменений) -->
        <!-- … -->
      -->

      <div class="max-w-md mx-auto bg-white dark:bg-gray-800 rounded-xl shadow-md overflow-hidden md:max-w-2xl mb-6">
        <div class="md:flex p-6">
          <div class="md:flex-shrink-0 flex items-center justify-center">
            <img
              v-if="user && user.avatar"
              :src="user.avatar"
              alt="User avatar"
              class="h-24 w-24 rounded-full object-cover"
            />
            <div v-else class="h-24 w-24 rounded-full bg-gray-300 dark:bg-gray-600 flex items-center justify-center text-gray-500 dark:text-gray-400 text-3xl font-semibold">
              {{ user ? user.name.charAt(0).toUpperCase() : '' }}
            </div>
          </div>
          <div class="mt-4 md:mt-0 md:ml-6 flex flex-col justify-center">
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white">{{ user ? user.name : '' }}</h2>
            <p class="text-gray-600 dark:text-gray-300">{{ user ? user.email : '' }}</p>
            <p class="text-gray-600 dark:text-gray-300 capitalize">{{ user ? user.role : '' }}</p>
            <button
              @click="logout"
              class="mt-4 px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg transition-colors duration-300"
            >
              Выйти
            </button>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Events Section -->
        <div class="lg:col-span-2">
          <div class="p-6 rounded-2xl events-section shadow-lg">
            <!-- Поиск и фильтры (без изменений) -->
            <!-- … -->

            <!-- Events List -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div
                v-for="event in filteredEvents"
                :key="event.id"
                class="p-4 rounded-xl event-card shadow-md cursor-pointer hover:shadow-lg transition-all duration-300"
                @click="goToEvent(event.id)"
              >
                <div class="flex justify-between items-start mb-2">
                  <h3 class="text-lg font-bold">{{ event.title }}</h3>
                  <span
                    class="px-2 py-1 rounded-full text-xs"
                    :class="{
                      'bg-blue-100 text-blue-800': event.status === 'upcoming',
                      'bg-green-100 text-green-800': event.status === 'current',
                      'bg-gray-100 text-gray-800': event.status === 'past'
                    }"
                  >
                    {{ event.status === 'upcoming' ? 'Предстоящее' :
                       event.status === 'current'  ? 'Текущее'     :
                       'Завершённое' }}
                  </span>
                </div>
                <p class="text-sm text-opacity-80 mb-2">{{ formatDate(event.date) }}</p>
                <p class="text-sm text-opacity-60">
                  {{ event.description ? event.description.substring(0, 100) + '...' : '' }}
                </p>
                <div class="mt-3 flex flex-wrap gap-2">
                  <span
                    v-for="(tag, i) in event.tags.split(',')"
                    :key="i"
                    class="px-2 py-1 rounded-full text-xs"
                    :class="tagColors[i % tagColors.length]"
                  >
                    {{ tag.trim() }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Column: Add New Event -->
        <div class="col-span-1 p-4 rounded-2xl add-event shadow-lg">
          <h2 class="text-xl font-semibold mb-4">Добавить мероприятие</h2>
          <form @submit.prevent="addEvent">
            <!-- Название -->
            <div class="mb-4 relative">
              <label class="block mb-1 text-opacity-80">Название</label>
              <input
                v-model="newEvent.title"
                type="text"
                class="w-full p-3 pl-10 rounded-xl input-field"
                placeholder="Название"
              />
            </div>

            <!-- Дата -->
            <div class="mb-4 relative">
              <label class="block mb-1 text-opacity-80">Дата (ДД.ММ.ГГГГ)</label>
              <input
                v-model="newEvent.date"
                type="text"
                class="w-full p-3 pl-10 rounded-xl input-field"
                placeholder="ДД.ММ.ГГГГ"
              />
            </div>

            <!-- Описание -->
            <div class="mb-4 relative">
              <label class="block mb-1 text-opacity-80">Описание</label>
              <textarea
                v-model="newEvent.description"
                class="w-full p-3 pl-10 rounded-xl input-field"
                placeholder="Описание"
                rows="3"
              ></textarea>
            </div>

            <!-- Теги -->
            <div class="mb-4 relative">
              <label class="block mb-1 text-opacity-80">Теги</label>
              <input
                v-model="newEvent.tags"
                type="text"
                class="w-full p-3 pl-10 rounded-xl input-field"
                placeholder="Теги через запятую"
              />
            </div>

            <!-- Файл изображения -->
            <div class="mb-4">
              <label class="block mb-1 text-opacity-80">Изображение</label>
              <input type="file" @change="onFileChange" class="w-full" />
              <p v-if="newEvent.image">{{ newEvent.image.name }}</p>
            </div>

            <button type="submit" class="p-3 rounded-xl action-btn">Добавить</button>
            <button type="button" class="p-3 rounded-xl cancel-btn" @click="resetForm">
              Отменить
            </button>
          </form>
        </div>
      </div>
    </div>

    <!-- Theme Toggle Button -->
    <button @click="toggleTheme" class="theme-toggle fixed bottom-4 right-4 p-3 rounded-full shadow-lg">
      <!-- … иконки … -->
    </button>
  </div>
</template>

<script>
import axios from 'axios'
import Navbar from '@/components/Navbar.vue'

export default {
  components: { Navbar },

  data() {
    return {
      user: null,
      error: null,

      // Чтобы filter/filterEvents всегда работал
      events: [],

      // 🛠️ Обязательно инициализируем newEvent, иначе newEvent.title будет undefined
      newEvent: {
        title: '',
        date: '',
        description: '',
        tags: '',
        image: null
      },

      isDarkMode: true,
      activeTab: 'participating',
      tabs: [
        { id: 'participating', label: 'Текущие' },
        { id: 'upcoming',      label: 'Предстоящие' },
        { id: 'participated',  label: 'Прошедшие' }
      ],
      tagColors: [
        'bg-blue-500/20 text-blue-500',
        'bg-purple-500/20 text-purple-500',
        'bg-green-500/20 text-green-500',
        'bg-yellow-500/20 text-yellow-500'
      ]
    }
  },

  async mounted() {
    axios.defaults.baseURL = 'http://localhost:8080'
    axios.defaults.withCredentials = true

    // Загружаем и профиль, и список мероприятий
    await Promise.all([ this.fetchUserData(), this.fetchEvents() ])
  },

  methods: {
    async fetchUserData() {
      try {
        const { data } = await axios.get('/profile')
        this.user = data
      } catch (e) {
        console.error('Error fetching user data:', e)
        this.error = 'Не удалось загрузить данные о пользователе.'
      }
    },
    // Заготовка под получение списка событий — чтобы filteredEvents работал
    async fetchEvents() {
      try {
        const { data } = await axios.get('/events') // реализуй этот маршрут на бэке
        this.events = data
      } catch (e) {
        console.warn('Не удалось загрузить мероприятия, используем локальные данные.')
        // можно оставить пустым — тогда events остаётся [] и ошибок нет
      }
    },
    logout() {
      axios.post('/logout')
      this.$router.push('/login')
    },
    toggleTheme() {
      this.isDarkMode = !this.isDarkMode
    },
    goToEvent(id) {
      this.$router.push(`/event/${id}`)
    },
    onFileChange(e) {
      this.newEvent.image = e.target.files[0]
    },
    addEvent() {
      // здесь можно собрать payload из this.newEvent и отправить POST /events
      console.log('Добавляем событие', this.newEvent)
      // После успешного добавления можно очистить форму:
      this.resetForm()
    },
    resetForm() {
      this.newEvent = { title: '', date: '', description: '', tags: '', image: null }
    },
    formatDate(d) {
      // форматиование строки даты
      return d
    }
  },

  computed: {
    // Чтобы не падало на undefined и всегда возвращать массив
    filteredEvents() {
      return (this.events || []).filter(evt => evt.status === this.activeTab)
    }
  }
}
</script>

<style scoped>
@import "leaflet/dist/leaflet.css";
/* Светлая тема */
.light-theme {
  background-color: #D5D5E0;
  color: #1A1A2E;
}

.light-theme .user-info,
.light-theme .add-event,
.light-theme .events-section {
  background-color: #E4E4F0;
}

.light-theme .event-card {
  background-color: #F0F0F5;
}

.light-theme .nav-btn,
.light-theme .action-btn,
.light-theme .theme-toggle {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.light-theme .logout-btn {
  background-color: #FF6584;
  color: #FFFFFF;
}

.light-theme .cancel-btn {
  background-color: #F0F0F5;
  color: #1A1A2E;
}

.light-theme .input-field,
.light-theme .file-input {
  background-color: #F0F0F5;
  border: 1px solid #E4E4F0;
  color: #1A1A2E;
}

.light-theme .tab-btn {
  background-color: #F0F0F5;
  color: #1A1A2E;
}

.light-theme .tab-active {
  background-color: #6C63FF;
  color: #FFFFFF;
}

/* Темная тема */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .user-info,
.dark-theme .add-event,
.dark-theme .events-section {
  background-color: #2A2A40;
}

.dark-theme .event-card {
  background-color: #3A3A50;
}

.dark-theme .nav-btn,
.dark-theme .action-btn,
.dark-theme .theme-toggle {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .logout-btn {
  background-color: #FF6584;
  color: #FFFFFF;
}

.dark-theme .cancel-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
}

.dark-theme .input-field,
.dark-theme .file-input {
  background-color: #3A3A50;
  border: 1px solid #E4E4F0;
  color: #E4E4F0;
}

.dark-theme .tab-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
}

.dark-theme .tab-active {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.map-placeholder {
  height: 100%;
  background-color: #ccc;
  display: flex;
  margin-bottom: 1%;
  align-items: center;
  justify-content: center;
}

.logo {
  color: #6C63FF;
}

.text-opacity-80 {
  opacity: 0.8;
}

.text-opacity-60 {
  opacity: 0.6;
}

.theme-toggle {
  transition: all 0.3s ease;
}

.theme-toggle:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}


#add-event-map {
  height: 100%;
  width: 100%;
  border-radius: 12px;
  z-index: 0;
}
.light-theme .nav-bar {
  background: rgba(228, 228, 240, 0.9);
  border-bottom: 1px solid rgba(26, 26, 46, 0.1);
}

.light-theme .nav-btn {
  background: rgba(240, 240, 245, 0.8);
  color: #1A1A2E;
}

.light-theme .nav-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
}

.light-theme .logout-btn {
  background: rgba(255, 101, 132, 0.8);
}

.light-theme .logout-btn:hover {
  background: #FF6584;
}

.light-theme .logo {
  color: #6C63FF;
}

.dark-theme .nav-bar {
  background: rgba(42, 42, 64, 0.9);
}

.dark-theme .nav-btn {
  background: rgba(58, 58, 80, 0.8);
  color: #E4E4F0;
}

.dark-theme .nav-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .logout-btn {
  background: rgba(255, 101, 132, 0.8);
}

.dark-theme .logout-btn:hover {
  background: #FF6584;
}

.dark-theme .logo {
  color: #6C63FF;
}
</style>
