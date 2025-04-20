<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']" style="padding-top: 64px;">
    <Navbar />
    
    <div class="container mx-auto p-4">
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
      <div class="flex flex-col md:flex-row gap-6 mb-6"> 
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3 flex-1">
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Мероприятий</p>
            <p class="text-2xl font-bold">24</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Друзей</p>
            <p class="text-2xl font-bold">56</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-green-500 to-teal-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Очков</p>
            <p class="text-2xl font-bold">1,240</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-yellow-500 to-orange-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Уровень</p>
            <p class="text-2xl font-bold">5</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-pink-500 to-red-500 rounded-full"></div>
          </div>
        </div>
      </div>

      

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Events Section -->
        <div class="lg:col-span-2">
          <div class="p-6 rounded-2xl events-section shadow-lg">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-xl font-bold">Мои мероприятия</h2>
              <button class="text-sm font-medium text-blue-500 hover:text-blue-400 transition-colors">
                Смотреть все
              </button>
            </div>
            

            <!-- Filters -->
            <div class="mb-6 flex flex-wrap gap-3">
              <button 
                v-for="(tab, index) in tabs" 
                :key="index"
                @click="activeTab = tab.id"
                :class="['px-4 py-2 rounded-lg transition-colors', activeTab === tab.id ? 'bg-blue-500 text-white' : 'bg-opacity-20 hover:bg-opacity-30']"
              >
                {{ tab.label }}
              </button>
            </div>

            <!-- Events List -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div
                v-for="event in filteredEvents"
                :key="event.id"
                class="p-4 rounded-xl event-card shadow-md cursor-pointer hover:shadow-lg transition-all duration-300 group"
                @click="goToEvent(event.id)"
              >
                <div class="relative h-40 rounded-lg overflow-hidden mb-3">
                  <img 
                    :src="event.image" 
                    class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
                    alt="Event image"
                  >
                  <div class="absolute inset-0 bg-gradient-to-t from-black/70 to-transparent"></div>
                  <div class="absolute bottom-3 left-3 right-3">
                    <h3 class="text-lg font-bold text-white">{{ event.title }}</h3>
                    <div class="flex justify-between items-center text-white text-opacity-80 text-sm">
                      <span>{{ event.date }}</span>
                      <span>{{ event.time }}</span>
                    </div>
                  </div>
                </div>
                <div class="flex justify-between items-center">
                  <div class="flex gap-1">
                    <span 
                      v-for="(tag, index) in event.tags.split(',')" 
                      :key="index"
                      class="px-2 py-1 text-xs rounded-full bg-opacity-20"
                      :class="tagColors[index % tagColors.length]"
                    >
                      {{ tag.trim() }}
                    </span>
                  </div>
                  <div class="flex items-center gap-1 text-sm">
                    <svg class="h-4 w-4" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                    </svg>
                    <span>{{ event.rating }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Sidebar -->
        <div class="space-y-6">
          <!-- Calendar -->
          <div class="p-6 rounded-2xl shadow-lg">
            <h3 class="text-lg font-bold mb-4">Календарь</h3>
            <div class="grid grid-cols-7 gap-1 text-center">
              <div class="text-sm font-medium py-2">Пн</div>
              <div class="text-sm font-medium py-2">Вт</div>
              <div class="text-sm font-medium py-2">Ср</div>
              <div class="text-sm font-medium py-2">Чт</div>
              <div class="text-sm font-medium py-2">Пт</div>
              <div class="text-sm font-medium py-2">Сб</div>
              <div class="text-sm font-medium py-2">Вс</div>
              
              <template v-for="(day, index) in calendarDays" :key="index">
                <div 
                  class="py-2 rounded-full"
                  :class="{
                    'bg-blue-500 text-white': day === currentDay,
                    'text-opacity-60': day > daysInMonth
                  }"
                >
                  {{ day > daysInMonth ? '' : day }}
                </div>
              </template>
            </div>
          </div>

          <!-- Friends Activity -->
          <div class="p-6 rounded-2xl shadow-lg">
            <h3 class="text-lg font-bold mb-4">Активность друзей</h3>
            <div class="space-y-4">
              <div 
                v-for="activity in friendActivities"
                :key="activity.id"
                class="flex items-center gap-3"
              >
                <div class="h-10 w-10 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-white text-sm font-bold">
                  {{ activity.initials }}
                </div>
                <div class="flex-1">
                  <p class="text-sm">
                    <span class="font-medium">{{ activity.name }}</span> {{ activity.action }}
                  </p>
                  <p class="text-xs text-opacity-60">{{ activity.time }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Theme Toggle -->
    <button @click="toggleTheme" class="theme-toggle fixed bottom-6 right-6 p-3 rounded-full shadow-lg hover:scale-110 transition-transform">
      <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
      </svg>
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
      events: [],
      isDarkMode: true,
      activeTab: 'participating',
      tabs: [
        { id: 'participating', label: 'Текущие' },
        { id: 'upcoming',    label: 'Предстоящие' },
        { id: 'participated', label: 'Прошедшие' }
      ],
      tagColors: [
        'bg-blue-500/20 text-blue-500',
        'bg-purple-500/20 text-purple-500',
        'bg-green-500/20 text-green-500',
        'bg-yellow-500/20 text-yellow-500'
      ],
    }
  },

  async mounted() {
    // Настраиваем axios перед первым запросом
    axios.defaults.baseURL = 'http://localhost:8080'
    axios.defaults.withCredentials = true

    await this.fetchUserData()
  },

  methods: {
    async fetchUserData() {
      try {
        const response = await axios.get('/profile')
        this.user = response.data
      } catch (e) {
        console.error('Error fetching user data:', e)
        this.error = 'Не удалось загрузить данные о пользователе.'
      }
    },
    logout() {
      // При logout удаляем токен на сервере / на клиенте
      // и возвращаемся на страницу логина
      axios.post('/logout')
      this.$router.push('/login')
    },
    toggleTheme() {
      this.isDarkMode = !this.isDarkMode
    },
    goToEvent(id) {
      this.$router.push(`/event/${id}`)
    }
  },

  computed: {
    filteredEvents() {
      // Защищённый .filter(), чтобы не было ошибки, если events ещё пуст
      return (this.events || []).filter(event => event.status === this.activeTab)
    }
  }
}
</script>
<style scoped>
/* Light Theme */
.light-theme {
  background-color: #f5f7fa;
  color: #1A1A2E;
}

.light-theme .user-info,
.light-theme .events-section,
.light-theme .stat-card {
  background-color: #ffffff;
}

.light-theme .event-card {
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
}

/* Dark Theme */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .user-info,
.dark-theme .events-section,
.dark-theme .stat-card {
  background-color: #2A2A40;
}

.dark-theme .event-card {
  background-color: #3A3A50;
  border: 1px solid #4A4A60;
}

/* Common Styles */
.user-info {
  transition: all 0.3s ease;
}

.stat-card {
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.event-card {
  transition: all 0.3s ease;
}

.event-card:hover {
  transform: translateY(-3px);
}

.theme-toggle {
  background-color: #6C63FF;
  color: white;
  transition: all 0.3s ease;
}

.text-opacity-80 {
  opacity: 0.8;
}

.text-opacity-60 {
  opacity: 0.6;
}

/* Animations */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.event-card {
  animation: fadeIn 0.5s ease forwards;
}

.event-card:nth-child(1) { animation-delay: 0.1s; }
.event-card:nth-child(2) { animation-delay: 0.2s; }
.event-card:nth-child(3) { animation-delay: 0.3s; }
.event-card:nth-child(4) { animation-delay: 0.4s; }
</style>
