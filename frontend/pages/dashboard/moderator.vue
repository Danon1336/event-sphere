<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']" style="padding-top: 64px;">
    <Navbar />
    
    <div class="container mx-auto p-4">
      <!-- Moderator Profile Header -->
      <div class="flex flex-col md:flex-row gap-6 mb-6">
        <!-- Profile Card -->
        <div class="p-6 rounded-2xl user-info shadow-lg flex-1 flex items-center gap-6">
          <div class="relative">
            <div class="h-20 w-20 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
              <span class="text-2xl font-bold text-white">ММ</span>
            </div>
            <div class="absolute -bottom-1 -right-1 h-6 w-6 rounded-full bg-yellow-500 border-2 border-white"></div>
          </div>
          <div class="flex-1">
            <h2 class="text-2xl font-bold">Мария Модераторова</h2>
            <p class="text-opacity-80 mb-2">Модератор платформы</p>
            <div class="flex gap-2">
              <span class="px-3 py-1 rounded-full text-sm bg-opacity-20 bg-blue-500">Senior Moderator</span>
              <span class="px-3 py-1 rounded-full text-sm bg-opacity-20 bg-green-500">Trusted</span>
            </div>
          </div>
        </div>

        <!-- Moderation Stats Cards -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3 flex-1">
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">На модерации</p>
            <p class="text-2xl font-bold">8</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Жалоб</p>
            <p class="text-2xl font-bold">12</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-red-500 to-pink-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Одобрено</p>
            <p class="text-2xl font-bold">42</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-green-500 to-teal-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Отклонено</p>
            <p class="text-2xl font-bold">5</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-yellow-500 to-orange-500 rounded-full"></div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Events Needing Moderation -->
        <div class="lg:col-span-2">
          <div class="p-6 rounded-2xl events-section shadow-lg">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-xl font-bold">Мероприятия на проверке</h2>
              <button class="text-sm font-medium text-blue-500 hover:text-blue-400 transition-colors">
                Смотреть все
              </button>
            </div>
            
            <!-- Moderation Filters -->
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
                  <div class="flex gap-2">
                    <button class="px-3 py-1 text-xs rounded-full bg-green-500/20 text-green-500 hover:bg-green-500/30">
                      Одобрить
                    </button>
                    <button class="px-3 py-1 text-xs rounded-full bg-red-500/20 text-red-500 hover:bg-red-500/30">
                      Отклонить
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Sidebar -->
        <div class="space-y-6">
          <!-- Reported Content -->
          <div class="p-6 rounded-2xl shadow-lg">
            <h3 class="text-lg font-bold mb-4">Жалобы пользователей</h3>
            <div class="space-y-4">
              <div 
                v-for="report in reports"
                :key="report.id"
                class="p-3 rounded-lg bg-opacity-20 hover:bg-opacity-30 transition-colors"
                :class="isDarkMode ? 'bg-red-500/10' : 'bg-red-500/10'"
              >
                <div class="flex items-center gap-3 mb-2">
                  <div class="h-8 w-8 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-white text-xs font-bold">
                    {{ report.userInitials }}
                  </div>
                  <div class="flex-1">
                    <p class="text-sm font-medium">{{ report.userName }}</p>
                    <p class="text-xs text-opacity-60">{{ report.time }}</p>
                  </div>
                </div>
                <p class="text-sm mb-2">{{ report.reason }}</p>
                <div class="flex gap-2">
                  <button class="px-2 py-1 text-xs rounded bg-blue-500/20 text-blue-500 hover:bg-blue-500/30">
                    Проверить
                  </button>
                  <button class="px-2 py-1 text-xs rounded bg-gray-500/20 text-gray-500 hover:bg-gray-500/30">
                    Отклонить
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Moderation Log -->
          <div class="p-6 rounded-2xl shadow-lg">
            <h3 class="text-lg font-bold mb-4">Журнал модерации</h3>
            <div class="space-y-3">
              <div 
                v-for="log in moderationLog"
                :key="log.id"
                class="flex items-start gap-3"
              >
                <div class="h-8 w-8 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-white text-xs font-bold mt-1">
                  {{ log.initials }}
                </div>
                <div class="flex-1">
                  <p class="text-sm">
                    <span class="font-medium">{{ log.action }}</span> - {{ log.target }}
                  </p>
                  <p class="text-xs text-opacity-60">{{ log.time }}</p>
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
      activeTab: 'pending',
      tabs: [
        { id: 'pending', label: 'Ожидают' },
        { id: 'reported', label: 'Жалобы' },
        { id: 'recent', label: 'Новые' }
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
          status: 'pending',
          image: 'https://source.unsplash.com/random/600x400/?conference',
        },
        { 
          id: 2, 
          title: 'Научный симпозиум', 
          date: '15 июня 2025', 
          time: '09:00 - 17:00',
          tags: 'наука, симпозиум, образование', 
          status: 'reported',
          image: 'https://source.unsplash.com/random/600x400/?science',
        },
        { 
          id: 3, 
          title: 'ТехноФорум 2025', 
          date: '20 августа 2025', 
          time: '11:00 - 19:00',
          tags: 'технологии, форум, инновации', 
          status: 'recent',
          image: 'https://source.unsplash.com/random/600x400/?tech',
        },
        { 
          id: 4, 
          title: 'Дизайн-уикенд', 
          date: '12 марта 2025', 
          time: '12:00 - 16:00',
          tags: 'дизайн, креатив', 
          status: 'pending',
          image: 'https://source.unsplash.com/random/600x400/?design',
        }
      ],
      reports: [
        { 
          id: 1, 
          userInitials: 'АП', 
          userName: 'Анна Петрова', 
          reason: 'Некорректное описание мероприятия', 
          time: '2 часа назад' 
        },
        { 
          id: 2, 
          userInitials: 'ИС', 
          userName: 'Иван Сидоров', 
          reason: 'Подозрительная активность организатора', 
          time: '5 часов назад' 
        },
        { 
          id: 3, 
          userInitials: 'МК', 
          userName: 'Мария Козлова', 
          reason: 'Несоответствие тематике', 
          time: 'вчера' 
        }
      ],
      moderationLog: [
        { 
          id: 1, 
          initials: 'ММ', 
          action: 'Одобрено мероприятие', 
          target: 'ИТ Конференция 2025', 
          time: 'Сегодня, 10:45' 
        },
        { 
          id: 2, 
          initials: 'ММ', 
          action: 'Отклонено мероприятие', 
          target: 'Нелегальный сбор', 
          time: 'Вчера, 18:30' 
        },
        { 
          id: 3, 
          initials: 'АС', 
          action: 'Заблокирован пользователь', 
          target: 'user123', 
          time: 'Вчера, 15:20' 
        }
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
  },
};
</script>

<style scoped>
/* Reuse the same styles from participant.vue */
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
