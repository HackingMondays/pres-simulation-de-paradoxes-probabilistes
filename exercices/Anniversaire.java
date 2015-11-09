import java.util.*;

/**
 * Created by rodrigo on 11/8/15.
 */
public class Anniversaire {

    static Random random = new Random(System.currentTimeMillis());

    static private double MAXVALUES = 365;

    static public List<Integer> getRandomValues(int count) {
        ArrayList<Integer> data = new ArrayList<>();
        for (int i=0; i<count; i++) {
            Double d = Math.floor(random.nextDouble() * MAXVALUES);
            data.add(d.intValue());
        }
        return data;
    }

    static public boolean hasDuplicates(List<Integer> data) {
        Set cache = new HashSet<Integer>();
        for (Iterator<Integer> i = data.iterator(); i.hasNext(); ) {
            Integer value = i.next();
            if (cache.contains(value)) {
                return true;
            }
            cache.add(value);
        }
        return false;
    }

    public static void main(String[] args) {
        System.out.println("ok");

        int populationSize = 23;
        int sampleSize = 100000;
        int success = 0;

        for (int i=0; i<sampleSize; i++) {
            List<Integer> data = getRandomValues(populationSize);
            if (hasDuplicates(data)) {
                ++success;
            }
        }

        System.out.println("Population size:" + populationSize);
        System.out.println("Iterations     :" + sampleSize);
        System.out.println("Same birthday  :" + success);
        System.out.println("Rate           :" + ((double)success/(double)sampleSize) * 100.0);
    }

}
