const HashTable = require('./HashTable')

describe('HashTable', () => {
    let hashTable;

    beforeEach(() => {
        hashTable = new HashTable(10);
    });

    test('should insert and retrieve a key-value pair', () => {
        hashTable.insert('key', 'value');
        expect(hashTable.get('key')).toBe('value');
    });

    test('should update value for existing key', () => {
        hashTable.insert('pill', 'blue');
        hashTable.insert('pill', 'red');
        expect(hashTable.get('pill')).toBe('red');
    });

    test('should handle hash collisions', () => {
        const key1 = 'rhic'; 
        const key2 = 'hadron'; 
        hashTable.insert(key1, 'fast');
        hashTable.insert(key2, 'faster');
        expect(hashTable.get(key1)).toBe('fast');
        expect(hashTable.get(key2)).toBe('faster');
    });

    test('should return null for non-existent key', () => {
        expect(hashTable.get('nonExistentKey')).toBeNull();
    });

    test('should return all key-value pairs', () => {
        hashTable.insert('key1', 'value1');
        hashTable.insert('key2', 'value2');
        console.log(hashTable.getAll())
        expect(hashTable.getAll()).toEqual([[], [], [], [], [], [], [], [], [['key1', 'value1']], [['key2', 'value2']]]);
    });

    test('should handle null and undefined keys gracefully', () => {
        hashTable.insert(null, 'nullValue');
        hashTable.insert(undefined, 'undefinedValue');
        expect(hashTable.get(null)).toBe('nullValue');
        expect(hashTable.get(undefined)).toBe('undefinedValue');
    });

    test('should update value for non-first key in collision list', () => {
        // Two different keys that hash to the same index
        const key1 = 'aliciaKeys'; 
        const key2 = 'keytar'; 

        hashTable.insert(key1, 'song');
        hashTable.insert(key2, 'music');
        hashTable.insert(key2, 'noise');

        expect(hashTable.get(key1)).toBe('song');
        expect(hashTable.get(key2)).toBe('noise');
    });

    test('should handle multiple collisions correctly', () => {
        // Multiple keys that all hash to the same index
        const key1 = 'masterKey'; 
        const key2 = 'kindaMasterKey';
        const key3 = 'reMasteredKey';

        hashTable.insert(key1, 'charizard');
        hashTable.insert(key2, 'charmeleon');
        hashTable.insert(key3, 'charmander');

        expect(hashTable.get(key1)).toBe('charizard');
        expect(hashTable.get(key2)).toBe('charmeleon');
        expect(hashTable.get(key3)).toBe('charmander');
    });

    test('should update value for a key not at the start or end of a collision chain', () => {
        const hashTable = new HashTable(5); // Small size so we force collisions

        // Choose keys that'll hash to the same index
        // These keys have character codes summing up a value with the same modulo size
        const key1 = 'a';  // Character code 97
        const key2 = 'f';  // Character code 102, 97 + 5 = 102
        const key3 = 'k';  // Character code 107, 102 + 5 = 107

        // We insert these keys in a way that forms a collision chain
        hashTable.insert(key1, 'obelisk'); 
        hashTable.insert(key2, 'wingedRa'); 
        hashTable.insert(key3, 'slifer'); 

        // Now we update the value of a key in the middle of the chain
        hashTable.insert(key2, 'exodia');

        expect(hashTable.get(key1)).toBe('obelisk');
        expect(hashTable.get(key2)).toBe('exodia');
        expect(hashTable.get(key3)).toBe('slifer');
    });

    test('should return null for a key that does not exist', () => {
        hashTable.insert('mario', 'redHat');

        // Attempt to get a value for a non-existing key
        const result = hashTable.get('waldo');

        expect(result).toBeNull();
    });
});
