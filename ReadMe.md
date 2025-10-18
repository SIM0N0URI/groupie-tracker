# Groupie Tracker

A web application that visualizes music artists and bands information by consuming and displaying data from a RESTful API.

## Overview

Groupie Tracker fetches data about various artists, their concert locations, dates, and relations between these data points to create an interactive user experience. The application demonstrates client-server communication and data visualization techniques.

## Features

- **Artist Information**: Display band/artist details including names, images, formation year, first album date, and members
- **Concert Locations**: Show last and upcoming concert venues
- **Concert Dates**: Display past and future concert dates
- **Relations**: Link artists with their concert dates and locations
- **Interactive Events**: Client-server communication through user-triggered actions
- **Data Visualization**: Multiple display formats (cards, tables, lists, etc.)

## API Structure

The application consumes four API endpoints:

1. **Artists** - Band/artist biographical information
2. **Locations** - Concert venue data
3. **Dates** - Concert schedule information
4. **Relation** - Connections between artists, dates, and locations

## Tech Stack

- **Backend**: Go (standard library only)
- **Frontend**: HTML/CSS
- **Data Format**: JSON

## Project Structure

```
groupie-tracker/
├── functions/
│   ├── About.go
│   ├── error.go
│   ├── Home.go
│   ├── style.go
│   └── utils.go
├── static/
│   ├── style.css
│   ├── styleArtist.css
│   └── styleError.css
├── templates/
│   ├── artist.html
│   ├── error.html
│   └── index.html
├── go.mod
└── main.go
```

## Installation

1. Clone the repository:
```bash
git clone https://learn.zone01oujda.ma/git/obenmbar/groupie-tracker.git
```

2. Navigate to the project directory:
```bash
cd groupie-tracker
```

3. Run the application:
```bash
go run main.go
```

4. Open your browser and visit `http://localhost:8080`

## Learning Objectives

- Data manipulation and storage
- JSON parsing and formatting
- HTML templating
- Event handling and display
- Client-server architecture
- RESTful API consumption
