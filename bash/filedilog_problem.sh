
# Check your current display
echo $DISPLAY

# Restart the portal with proper display
systemctl --user stop xdg-desktop-portal-gtk
systemctl --user stop xdg-desktop-portal

# Make sure DISPLAY is set in systemd user environment
systemctl --user import-environment DISPLAY WAYLAND_DISPLAY XDG_CURRENT_DESKTOP

# Start them again
systemctl --user start xdg-desktop-portal
