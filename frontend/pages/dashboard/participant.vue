<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']" style="padding-top: 64px;">
    <Navbar />
    
    <div class="container mx-auto p-4">
      <!-- User Profile Header -->
      <div class="flex flex-col md:flex-row gap-6 mb-6">
        <!-- Profile Card -->
        <div class="p-6 rounded-2xl user-info shadow-lg flex-1 flex items-center gap-6">
          <div class="relative">
            <div class="h-20 w-20 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
              <span class="text-2xl font-bold text-white">ИИ</span>
            </div>
            <div class="absolute -bottom-1 -right-1 h-6 w-6 rounded-full bg-green-500 border-2 border-white"></div>
          </div>
          <div class="flex-1">
            <h2 class="text-2xl font-bold">Иван Иванов</h2>
            <p class="text-opacity-80 mb-2">Участник мероприятий</p>
            <div class="flex gap-2">
              <span class="px-3 py-1 rounded-full text-sm bg-opacity-20 bg-blue-500">Gold Member</span>
              <span class="px-3 py-1 rounded-full text-sm bg-opacity-20 bg-purple-500">Top 10%</span>
            </div>
          </div>
        </div>

        <!-- Stats Cards -->
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
export default {
  data() {
    const now = new Date();
    const firstDay = new Date(now.getFullYear(), now.getMonth(), 1).getDay();
    const daysInMonth = new Date(now.getFullYear(), now.getMonth() + 1, 0).getDate();
    
    return {
      isDarkMode: true,
      activeTab: 'participating',
      tabs: [
        { id: 'participating', label: 'Текущие' },
        { id: 'upcoming', label: 'Предстоящие' },
        { id: 'participated', label: 'Прошедшие' }
      ],
      currentDay: now.getDate(),
      daysInMonth,
      calendarDays: Array.from({ length: 42 }, (_, i) => i - firstDay + 2),
      tagColors: [
        'bg-blue-500/20 text-blue-500',
        'bg-purple-500/20 text-purple-500',
        'bg-green-500/20 text-green-500',
        'bg-yellow-500/20 text-yellow-500'
      ],
      events: [
        { 
          id: 1, 
          title: 'ИТ Конференция 2025', 
          date: '10 апреля 2025', 
          time: '10:00 - 18:00',
          tags: 'конференция, ИТ, технологии', 
          status: 'participating',
          image: 'https://source.unsplash.com/random/600x400/?conference',
          rating: 4.8
        },
        { 
          id: 2, 
          title: 'Научный симпозиум', 
          date: '15 июня 2025', 
          time: '09:00 - 17:00',
          tags: 'наука, симпозиум, образование', 
          status: 'upcoming',
          image: 'https://source.unsplash.com/random/600x400/?science',
          rating: 4.5
        },
        { 
          id: 3, 
          title: 'ТехноФорум 2025', 
          date: '20 августа 2025', 
          time: '11:00 - 19:00',
          tags: 'технологии, форум, инновации', 
          status: 'upcoming',
          image: 'https://source.unsplash.com/random/600x400/?tech',
          rating: 4.7
        },
        { 
          id: 4, 
          title: 'Дизайн-уикенд', 
          date: '12 марта 2025', 
          time: '12:00 - 16:00',
          tags: 'дизайн, креатив', 
          status: 'participated',
          image: 'https://source.unsplash.com/random/600x400/?design',
          rating: 4.9
        }
      ],
      friendActivities: [
        { id: 1, initials: 'АП', name: 'Анна Петрова', action: 'зарегистрировалась на ИТ Конференцию', time: '2 часа назад' },
        { id: 2, initials: 'ИС', name: 'Иван Сидоров', action: 'получил достижение "Активный участник"', time: '5 часов назад' },
        { id: 3, initials: 'МК', name: 'Мария Козлова', action: 'добавила новое мероприятие', time: 'вчера' }
      ]
    };
  },
  computed: {
    filteredEvents() {
      return this.events.filter(event => event.status === this.activeTab).slice(0, 4);
    },
  },
  methods: {
    toggleTheme() {
      this.isDarkMode = !this.isDarkMode;
    },
    goToEvent(id) {
      this.$router.push(`/event/${id}`);
    },
  },
};
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
