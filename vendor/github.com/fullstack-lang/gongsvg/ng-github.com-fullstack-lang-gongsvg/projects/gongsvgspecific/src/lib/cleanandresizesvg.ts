export function processSVG(svgString: string): string {
    // Parse the SVG string into a DOM document
    const parser = new DOMParser();
    const doc = parser.parseFromString(svgString, "image/svg+xml");
    const svg = doc.documentElement;
  
    // Find all rect elements
    const rects = svg.getElementsByTagName('rect');
    
    // Remove the first rect if it exists
    if (rects.length > 0) {
      rects[0].remove();
    }
  
    // Get all remaining elements and convert to array for iteration
    const elements = Array.from(svg.getElementsByTagName('*'));
    
    // Initialize boundaries
    let minX = Infinity;
    let minY = Infinity;
    let maxX = -Infinity;
    let maxY = -Infinity;
  
    // Process each element to find boundaries
    for (const element of elements) {
      // Get element's bounding box if available
      try {
        const bbox = (element as SVGGraphicsElement).getBBox?.();
        if (bbox) {
          // Consider transform attributes
          const transform = element.getAttribute('transform');
          let tx = 0, ty = 0;
          
          if (transform) {
            const match = transform.match(/translate\(([-\d.]+)\s*([-\d.]+)?\)/);
            if (match) {
              tx = parseFloat(match[1]) || 0;
              ty = parseFloat(match[2]) || 0;
            }
          }
  
          // Update boundaries considering transforms
          minX = Math.min(minX, bbox.x + tx);
          minY = Math.min(minY, bbox.y + ty);
          maxX = Math.max(maxX, bbox.x + bbox.width + tx);
          maxY = Math.max(maxY, bbox.y + bbox.height + ty);
        }
      } catch (e) {
        // Some elements might not support getBBox
        console.warn('Unable to get bounding box for element:', element);
      }
  
      // Handle explicit coordinates for elements like lines
      ['x', 'x1', 'x2', 'cx'].forEach(attr => {
        const val = parseFloat(element.getAttribute(attr) || '');
        if (!isNaN(val)) {
          minX = Math.min(minX, val);
          maxX = Math.max(maxX, val);
        }
      });
  
      ['y', 'y1', 'y2', 'cy'].forEach(attr => {
        const val = parseFloat(element.getAttribute(attr) || '');
        if (!isNaN(val)) {
          minY = Math.min(minY, val);
          maxY = Math.max(maxY, val);
        }
      });
    }
  
    // Add padding
    const padding = 10;
    minX -= padding;
    minY -= padding;
    maxX += padding;
    maxY += padding;
  
    // Calculate new dimensions
    const width = maxX - minX;
    const height = maxY - minY;
  
    // Update SVG attributes
    svg.setAttribute('width', width.toString());
    svg.setAttribute('height', height.toString());
    svg.setAttribute('viewBox', `${minX} ${minY} ${width} ${height}`);
  
    // Convert back to string
    return new XMLSerializer().serializeToString(svg);
  }